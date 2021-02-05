package cc.jbx.s3tester;

import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;

import software.amazon.awssdk.core.sync.RequestBody;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.PutObjectRequest;
import software.amazon.awssdk.services.s3.model.PutObjectResponse;
import software.amazon.awssdk.services.s3.model.S3Exception;

public class Main {
    public static void main(String[] args) {
        S3Client s3 = S3Client.builder()
            .region(Region.US_EAST_1)
            .build();

        String bucketName = System.getenv("BUCKET_NAME");
        try {
            Map<String, String> metadata = new HashMap<>();
            PutObjectRequest putOb = PutObjectRequest.builder()
                .bucket(bucketName)
                .key("javatest")
                .metadata(metadata)
                .build();

            byte[] b = "hello world".getBytes(StandardCharsets.UTF_8);
            PutObjectResponse response = s3.putObject(putOb, RequestBody.fromBytes(b));
            System.out.println(response.eTag());
        } catch (S3Exception e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}