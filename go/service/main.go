package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/pwmcintyre/logging/go/logger"
)

type Action int

const (
	ThingWrite Action = iota
	ThingRead
)

type Authorizer interface {
	ActionAuthorized(ctx context.Context, a Action) (bool, error)
}

type Thing struct {
	ID   string
	Body interface{}
}
type ThingPutter interface {
	Save(ctx context.Context, t Thing) error
}

var ErrSystem = errors.New("system error")
var ErrClient = errors.New("client error")

type SaveThingRequest struct {
	Thing Thing
}

type ThingService struct {
	auth  Authorizer
	store ThingPutter
}

func (s *ThingService) SaveThing(ctx context.Context, r SaveThingRequest) error {

	// auth
	ok, err := s.auth.ActionAuthorized(ctx, ThingWrite)
	if err != nil {
		// observe system failure
		logrus.WithError(err).Error("failed to authenticate")

		// return generic error so as to not leak internals
		return ErrSystem
	}
	if !ok {
		return errors.Wrap(ErrClient, "unauthorized")
	}

	// save
	if err := s.store.Save(ctx, r.Thing); err != nil {
		// observe system failure
		logger.WithContext(ctx).WithError(err).Error("failed to save thing")

		// return generic error so as to not leak internals
		return ErrSystem
	}

	// observe success
	logger.
		WithContext(ctx).
		WithField("thing_id", r.Thing.ID).
		Info("things saved")

	return nil

}

type Observer interface {
	AuthError(ctx context.Context, r SaveThingRequest, err error)
	Unauthorized(ctx context.Context, r SaveThingRequest)
	SaveError(ctx context.Context, r SaveThingRequest, err error)
	Save(ctx context.Context, r SaveThingRequest)
}

type ThingServiceWithObserver struct {
	auth    Authorizer
	store   ThingPutter
	observe Observer
}

func (s *ThingServiceWithObserver) SaveThing(ctx context.Context, r SaveThingRequest) error {

	// auth
	ok, err := s.auth.ActionAuthorized(ctx, ThingWrite)
	if err != nil {
		s.observe.AuthError(ctx, r, err)
	}
	if !ok {
		s.observe.Unauthorized(ctx, r)
		return errors.Wrap(ErrClient, "unauthorized")
	}

	// save
	if err := s.store.Save(ctx, r.Thing); err != nil {
		s.observe.SaveError(ctx, r, err)
	}

	// emit
	s.observe.Save(ctx, r)
	return nil

}

type Tracer interface {
	Auth(ctx context.Context, fn func() error) error
	Save(ctx context.Context, fn func() error) error
}

type LogTracer struct{}

func (t *LogTracer) Save(ctx context.Context, thing Thing, fn func() error) error {
	start := time.Now()
	err := fn()
	l := logger.
		WithContext(ctx).
		WithField("duration", time.Since(start)).
		WithField("thing_id", thing.ID)
	if err != nil {
		l.Error("failed to save thing")
		return err
	}
	l.Info("things saved")
	return nil
}

type ThingServiceWithTracer struct {
	auth  Authorizer
	store ThingPutter
	trace Tracer
}

func (s *ThingServiceWithTracer) SaveThing(ctx context.Context, r SaveThingRequest) error {

	// auth
	var ok bool
	err := s.trace.Auth(ctx, func() (err error) {
		ok, err = s.auth.ActionAuthorized(ctx, ThingWrite)
		return
	})
	if err != nil {
		return ErrSystem
	}
	if !ok {
		return errors.Wrap(ErrClient, "unauthorized")
	}

	// save
	if s.trace.Save(ctx, func() error {
		return s.store.Save(ctx, r.Thing)
	}); err != nil {
		return ErrSystem
	}

	// emit
	return nil

}
