package org.rcmendes.echo.core;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;


@Path("/status")
public class StatusResource {

    @Inject
    EchoService service;

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    @Path("/")
    public Object status() {
        var message = "{\"status\":\"OK\"}";

        return message;
    }
}