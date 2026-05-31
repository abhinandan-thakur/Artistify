package com.abhi.notification_service.grpc;

import org.springframework.stereotype.Service;

import com.abhi.notification_service.dto.EmailRequest;
import com.abhi.notification_service.service.EmailProducer;

import io.grpc.stub.StreamObserver;
import mailing.Mailing;
import mailing.MailingServiceGrpc;

@Service
public class MailingGrpcService
        extends MailingServiceGrpc
        .MailingServiceImplBase {

    private final EmailProducer emailProducer;

    public MailingGrpcService(
            EmailProducer emailProducer
    ) {
        this.emailProducer =
                emailProducer;
    }

    @Override
    public void sendMail(
            Mailing.MailingServiceRequest request,

            StreamObserver<
                    Mailing.MailingServiceResponse
            > responseObserver
    ) {

        System.out.println(
                "Received mail request for: "
                + request.getReceiverEmail()
        );

        EmailRequest email =
                new EmailRequest();

        email.setReceiverEmail(
                request.getReceiverEmail()
        );

        email.setOtp(
                request.getOtp()
        );

        emailProducer.queueEmail(
                email
        );

Mailing.MailingServiceResponse
        response =
        Mailing.MailingServiceResponse
                .newBuilder()
                .setSucces(
                        true
                )
                .setResponse(
                        "Email queued"
                )
                .build();

        responseObserver.onNext(
                response
        );

        responseObserver.onCompleted();
    }
}