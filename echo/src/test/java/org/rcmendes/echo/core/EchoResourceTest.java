package org.rcmendes.echo.core;

import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;
import static org.hamcrest.CoreMatchers.is;


@QuarkusTest
public class EchoResourceTest {

    @Test
    public void testEchoEndpoint() {
        String entry = "This is a Test";
        given().pathParam("name", entry).when().get("/echo/shout/{name}").then().statusCode(200)
                .body(is("{\"message\":\"THIS IS A TEST\"}"));
    }

}