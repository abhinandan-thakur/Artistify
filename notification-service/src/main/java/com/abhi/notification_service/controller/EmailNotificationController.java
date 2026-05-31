package com.abhi.notification_service.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.abhi.notification_service.dto.EmailRequest;
import com.abhi.notification_service.service.EmailProducer;

@RestController
@RequestMapping("/api")
public class EmailNotificationController {

    private final EmailProducer emailProducer;

    public EmailNotificationController(EmailProducer emailPEmailProducer) {
        this.emailProducer = emailPEmailProducer;
    }

    @PostMapping("/email")
    public String emailNotification(@RequestBody EmailRequest request) {
        emailProducer.queueEmail(request);
        return "Email queued";
    }
}