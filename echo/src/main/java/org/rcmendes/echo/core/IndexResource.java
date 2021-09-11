package org.rcmendes.echo.core;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;


@Path("/")
public class IndexResource {

    @Inject
    EchoService service;

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    @Path("/")
    public Object index() {
        var message = "Welcome to Echo API!";
        return String.format("{\"message\":\"%s\"", message);
    }
}