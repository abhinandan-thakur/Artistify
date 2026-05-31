package com.abhi.notification_service.service;

import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;

import com.abhi.notification_service.dto.EmailRequest;

@Service
public class EmailConsumer {
    private final RedisTemplate<String, Object> redisTemplate;
    private final JavaMailSender mailSender;
    private final EmailService emailService;

    public EmailConsumer(RedisTemplate<String,Object> redisTemplate, JavaMailSender mailSender, EmailService emailService) {
        this.redisTemplate = redisTemplate;
        this.mailSender = mailSender;
        this.emailService = emailService;
    }

    @Scheduled(fixedRate = 3000)
    public void consumeQueue() {
        
        EmailRequest emailRequest = (EmailRequest) redisTemplate.opsForList().leftPop("email_queue");
        if (emailRequest == null) {
            return;
        }
        emailRequest =
        emailService.prepareEmail(
                emailRequest
        );
        SimpleMailMessage mail = new SimpleMailMessage();
        mail.setTo(emailRequest.getReceiverEmail());
        mail.setFrom(emailRequest.getSenderEmail());
        mail.setSubject(emailRequest.getSubject());
        mail.setText(emailRequest.getBody());
        mailSender.send(mail);

        System.out.println("Mail sent to:"+emailRequest.getReceiverEmail());
    }

}