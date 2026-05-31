package com.abhi.notification_service.service;

import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import com.abhi.notification_service.dto.EmailRequest;

@Service
public class EmailProducer {
    private final RedisTemplate<String, Object> redisTemplate;

    public EmailProducer(RedisTemplate<String, Object> redisTemplate) {
        this.redisTemplate = redisTemplate;
    }

    public void queueEmail(EmailRequest emailRequest) {
        redisTemplate.opsForList().rightPush("email_queue", emailRequest);
    }
}
