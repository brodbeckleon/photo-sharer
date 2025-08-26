package brodbleo.friends.photo_sharer.auth;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@RestController
@RequestMapping("/api/auth")
public class AuthController {

    private final JwtService jwtService;
    private final String loginCode;

    public AuthController(JwtService jwtService, @Value("${app.login-code}") String loginCode) {
        this.jwtService = jwtService;
        this.loginCode = loginCode;
    }

    @PostMapping("/login")
    public ResponseEntity<Map<String, String>> login(@RequestBody LoginRequest loginRequest) {
        if (loginCode.equals(loginRequest.code())) {
            String token = jwtService.generateToken("user");
            return ResponseEntity.ok(Map.of("token", token));
        }
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED).build();
    }

    // DTO for the request body
    record LoginRequest(String code) {}
}