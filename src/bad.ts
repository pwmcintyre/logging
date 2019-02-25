import { S3 } from "aws-sdk";

export class Bad {

    constructor(
        public client = new S3()
    ){}
    
    public async handle (correlationId:string) {
    
        try {

            console.log( "correlationId=" + correlationId )

            const count = await this.getBucketCount()
            console.log( `found ${ count } buckets` )

            await this.makeBucket( correlationId )
            console.log( `bucket created` )
    
        } catch (error) {
            console.log( error )
        }
    
    }

    private async getBucketCount () {
        const buckets = await this.client.listBuckets().promise()
        return buckets.Buckets.length
    }

    private async makeBucket (Bucket: string) {
        await this.client.createBucket({Bucket}).promise()
        console.log( "created" )
    }

}
