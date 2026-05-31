package com.abhi.notification_service.dto;

import java.io.Serializable;

public class EmailRequest
        implements Serializable {

    private String senderEmail;
    private String receiverEmail;
    private String receiverUsername;
    private String subject;
    private String body;
    private String otp;

    public EmailRequest() {}

    public String getSenderEmail() {
        return senderEmail;
    }

    public void setSenderEmail(
            String senderEmail
    ) {
        this.senderEmail =
                senderEmail;
    }

    public String getReceiverEmail() {
        return receiverEmail;
    }

    public void setReceiverEmail(
            String receiverEmail
    ) {
        this.receiverEmail =
                receiverEmail;
    }

    public String getReceiverUsername() {
        return receiverUsername;
    }

    public void setReceiverUsername(
            String receiverUsername
    ) {
        this.receiverUsername =
                receiverUsername;
    }

    public String getSubject() {
        return subject;
    }

    public void setSubject(
            String subject
    ) {
        this.subject = subject;
    }

    public String getBody() {
        return body;
    }

    public void setBody(
            String body
    ) {
        this.body = body;
    }

    public String getOtp() {
        return otp;
    }

    public void setOtp(
            String otp
    ) {
        this.otp = otp;
    }
}