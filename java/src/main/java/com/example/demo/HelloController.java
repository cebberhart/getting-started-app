package com.example.demo;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.LinkedHashMap;
import java.util.Map;

@RestController
public class HelloController {

    @GetMapping("/")
    public Map<String, String> root() {
        Map<String, String> body = new LinkedHashMap<>();
        body.put("language", "Java (Spring Boot)");
        body.put("message", "Hello from a containerized Java app!");
        body.put("status", "running");
        return body;
    }

    @GetMapping("/health")
    public Map<String, String> health() {
        return Map.of("status", "ok");
    }
}
