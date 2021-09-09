package org.rcmendes.echo.core;

import javax.enterprise.context.ApplicationScoped;

@ApplicationScoped
public class EchoService {

    public String echo(String value) {
        return value.toUpperCase();
    }

}
