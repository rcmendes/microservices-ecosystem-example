package org.rcmendes.echo.core;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import org.jboss.resteasy.annotations.jaxrs.PathParam;

@Path("/echo")
public class EchoResource {

    @Inject
    EchoService service;

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    @Path("/shout/{name}")
    public Object shout(@PathParam String name) {
        var shout = service.echo(name);
        var message = "{\"message\":\"" + shout + "\"}";

        return message;
    }
}