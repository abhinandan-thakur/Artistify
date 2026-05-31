package com.abhi.notification_service.service;

import java.security.SecureRandom;

import org.springframework.stereotype.Service;

import com.abhi.notification_service.dto.EmailRequest;

@Service
public class EmailService {

    public EmailRequest prepareEmail(EmailRequest request) {

        String otp = generateOTP(6);

        request.setSenderEmail("celebchroniclesclient@gmail.com");

        request.setSubject("Login Authentication OTP Verification");

        request.setBody(
                "Hi "
                + request.getReceiverUsername()
                + ", your OTP verification for email "
                + request.getReceiverEmail()
                + " is: "
                + otp
        );

        return request;
    }

    private String generateOTP(
            int digits
    ) {

        SecureRandom random = new SecureRandom();

        StringBuilder otp = new StringBuilder();

        for (
                int i = 0;
                i < digits;
                i++
        ) {
            otp.append(
                    random.nextInt(10)
            );
        }

        return otp.toString();
    }
}