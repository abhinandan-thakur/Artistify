package com.abhi.notification_service.dto;

public class EmailResponse {

    private String receiverUsername;
    private String receiverEmail;
    private String senderEmail;
    private String subject;
    private String body;
    private String otp;
    private String fullResponse;

    public EmailResponse() {}

    public EmailResponse(
            String receiverUsername,
            String receiverEmail,
            String senderEmail,
            String subject
    ) {
        this.receiverUsername =
                receiverUsername;
        this.receiverEmail =
                receiverEmail;
        this.senderEmail =
                senderEmail;
        this.subject = subject;
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

    public String getReceiverEmail() {
        return receiverEmail;
    }

    public void setReceiverEmail(
            String receiverEmail
    ) {
        this.receiverEmail =
                receiverEmail;
    }

    public String getSenderEmail() {
        return senderEmail;
    }

    public void setSenderEmail(
            String senderEmail
    ) {
        this.senderEmail =
                senderEmail;
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

    public String getFullResponse() {
        return fullResponse;
    }

    public void setFullResponse(
            String fullResponse
    ) {
        this.fullResponse =
                fullResponse;
    }
}