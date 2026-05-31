package com.abhi.notification_service.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.abhi.notification_service.grpc.MailingGrpcService;

import io.grpc.Server;
import io.grpc.ServerBuilder;

@Configuration
public class GrpcServerConfig {

    @Bean(
            initMethod = "start",
            destroyMethod = "shutdown"
    )
    public Server grpcServer(
            MailingGrpcService service
    ) {

        return ServerBuilder
                .forPort(9090)
                .addService(service)
                .build();
    }
}