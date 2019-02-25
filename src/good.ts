import { S3 } from "aws-sdk";
import * as logger from "lambda-log";

// filtering logs:
// cat good.json | jq '. | select( .correlationId == "01D499VS1T55AWCV9PGH7AQYS3" )'
// cat good.json | jq '. | select( ._logLevel == "error" ) | .error.code'

export class Good {

    constructor(
        public client = new S3()
    ){}

    public async handle (correlationId:string) {

        try {

            logger.info( "start", { correlationId} )

            const count = await this.getBucketCount()
            logger.info( "buckets", { correlationId, count } )

            await this.makeBucket( correlationId )
            logger.info( "created", { correlationId } )

        } catch (error) {
            logger.error( "failure", { correlationId, error } )
        }

    }

    private async getBucketCount () {
        const buckets = await this.client.listBuckets().promise()
        return buckets.Buckets.length
    }

    private async makeBucket (Bucket: string) {
        await this.client.createBucket({Bucket}).promise()
    }

}
