package cc.jbx.s3tester;

import java.util.List;

import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;
import com.amazonaws.services.s3.model.Bucket;
import com.amazonaws.services.s3.model.PutObjectResult;

public class Main {
    public static void main(String[] args) {
        String bucketName = System.getenv("BUCKET_NAME");

        AmazonS3ClientBuilder clientBuilder = AmazonS3ClientBuilder.standard();
        AmazonS3 svc = clientBuilder.build();
        List<Bucket> buckets = svc.listBuckets();

        if (buckets.size() == 0) {
            System.err.println("No buckets found!");
            System.exit(1);
        }

        for (Bucket b : buckets) {
            System.out.println(b.getName());
        }

        PutObjectResult resp = svc.putObject(bucketName, "javatest", "hello world");
        System.out.println(resp.getETag());
    }
}