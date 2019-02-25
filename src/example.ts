import { ulid } from "ulid";
import { Bad } from "./bad";
import { Good } from "./good";

run(10)

export function run (count: number) {

    const good = new Good()
    const bad = new Bad()

    const i = setInterval( ()=> {

        // when to stop
        if ( count-- > 0 ) {

            // run
            bad.handle( ulid() )
            // good.handle( ulid() )

        } else {
            clearInterval( i )
        }

    }, 100 )

}
