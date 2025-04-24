package com.student.portal.results;

import org.springframework.web.bind.annotation.*;

import java.util.*;

@RestController
@RequestMapping("/results")
public class ResultController {

    private final List<String> dummyResults = new ArrayList<>(List.of("Math - A", "Science - B+"));

    @GetMapping
    public List<String> getResults() {
        return dummyResults;
    }

    @PostMapping
    public List<String> addResult(@RequestBody String result) {
        dummyResults.add(result);
        return dummyResults;
    }
}
