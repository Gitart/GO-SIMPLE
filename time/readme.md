# How to format time in Go/Golang

Go uses a special "magic" reference time that might seem weird at first:
The Magic Reference Time is:
Or put another way:
January 2, 2006 at 3:04:05 PM MST
Here's the genius part - the numbers in this date line up in order:

Month: 1 
Day: 2  
Hour: 3 
Minute: 4  
Second: 5 
Year: 6 
Let me show you with a super simple cheat sheet:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    // SUPER EASY CHEAT SHEET!

    // Just remember: 1 2 3 4 5 6
    fmt.Println(now.Format("1"))                    // => 1 (month)
    fmt.Println(now.Format("2"))                    // => 2 (day)
    fmt.Println(now.Format("3"))                    // => 3 (hour)
    fmt.Println(now.Format("4"))                    // => 4 (minute)
    fmt.Println(now.Format("5"))                    // => 5 (second)
    fmt.Println(now.Format("6"))                    // => 6 (year)

    // Real-world examples:

    // Want date like "2024-03-27"?
    fmt.Println(now.Format("2006-01-02"))

    // Want time like "15:04"?
    fmt.Println(now.Format("15:04"))

    // Want something like "Mar 27, 2024 3:04pm"?
    fmt.Println(now.Format("Jan 2, 2006 3:04pm"))

    // Want milliseconds? Add .000
    fmt.Println(now.Format("15:04:05.000"))

    // Common formats people use:
    formats := map[string]string{
        "Simple date":       "2006-01-02",          // => 2024-03-27
        "Simple time":       "15:04",               // => 13:45
        "Date and time":     "2006-01-02 15:04:05", // => 2024-03-27 13:45:30
        "Pretty date":       "January 2, 2006",     // => March 27, 2024
        "Short date":        "Jan 2",               // => Mar 27
        "Kitchen time":      "3:04PM",              // => 1:45PM
        "File safe":         "20060102-150405",     // => 20240327-134530
    }

    // Print them all
    for name, format := range formats {
        fmt.Printf("%s: %s\n", name, now.Format(format))
    }
}
```

            
# The trick to remember:

It's all based on one specific time: January 2, 2006 at 3:04:05 PM
Just write the date/time exactly how you want it to look, but use the reference time's numbers
Some easy examples:

```go
// Need just the month and day?
"01-02"                 // => "03-27"

// Need year, month, day?
"2006-01-02"            // => "2024-03-27"

// Need hours and minutes in 24h format?
"15:04"                 // => "13:45"

// Need everything?
"2006-01-02 15:04:05"   // => "2024-03-27 13:45:30"
```
                
## Pro Tips:

Need 24-hour time? Use "15" for hours
Need 12-hour time? Use "3" for hours
Need PM/AM? Just write "PM" or "pm" where you want it
Need month name? Use "January" or "Jan"
Think of it like a template - you're just showing Go how you want the date to look, using that special reference date as your model!

Now let's see how to handle timezones:
If we don't pay attention to the timezones, we can easily get wrong results.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // WRONG WAY (dangerous on servers!):
    localTime := time.Now() // Don't use this alone on servers!

    // RIGHT WAY (safe for servers):
    utcTime := time.Now().UTC()

    // Let's see different ways to handle timezones:

    // 1. Always store in UTC (RECOMMENDED FOR SERVERS)
    now := time.Now().UTC()

    // Format with UTC explicitly
    fmt.Println(now.Format("2006-01-02 15:04:05 MST"))   // Shows UTC
    fmt.Println(now.Format("2006-01-02 15:04:05 -0700")) // Shows +0000
    fmt.Println(now.Format(time.RFC3339))                // ISO8601/RFC3339 format

    // 2. Loading specific timezones (when you need to)
    nyc, _ := time.LoadLocation("America/New_York")
    baku, _ := time.LoadLocation("Asia/Baku")

    // Convert UTC to specific timezone
    nycTime := now.In(nyc)
    bakuTime := now.In(baku)

    fmt.Printf("UTC:   %s\n", now.Format(time.RFC3339))
    fmt.Printf("NYC:   %s\n", nycTime.Format(time.RFC3339))
    fmt.Printf("Baku: %s\n", bakuTime.Format(time.RFC3339))

    // 3. Parsing times with timezones
    // Always parse assuming UTC unless timezone is specified
    const format = "2006-01-02 15:04:05"

    // SAFE way to parse times without timezone
    timeStr := "2024-03-27 15:04:05"
    parsedTime, _ := time.Parse(format, timeStr)         // Assumes UTC

    // SAFE way to parse times with specific timezone
    parsedWithZone, _ := time.ParseInLocation(format, timeStr, nyc)

    // 4. Common patterns for server applications
    serverPatterns := map[string]string{
        // Store this format in database (UTC):
        "database": "2006-01-02 15:04:05",

        // Use this for JSON/API responses (includes timezone):
        "api": time.RFC3339,

        // Use this for logs:
        "logs": "2006-01-02 15:04:05 -0700",
    }

    for name, pattern := range serverPatterns {
        fmt.Printf("%s: %s\n", name, now.Format(pattern))
    }
}
```
        
Here are the GOLDEN RULES for server timezone handling:

## 1. ALWAYS Store in UTC:

```go
// RIGHT WAY - Store this in your database
timestamp := time.Now().UTC()

// WRONG WAY - Don't store local time
timestamp := time.Now() // Dangerous! Server timezone might change!
```

## 2. ALWAYS Parse Without Timezone as UTC:

```go
// RIGHT WAY - Parse times without timezone
timeStr := "2024-03-27 15:04:05"
safeTime, _ := time.Parse("2006-01-02 15:04:05", timeStr) // Assumes UTC

// WRONG WAY - Don't use ParseInLocation unless you specifically need a timezone
dangerousTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
```
        
## 3. ALWAYS Use RFC3339 for APIs:

```go
// RIGHT WAY - Use RFC3339 for API responses
type APIResponse struct {
    Timestamp string `json:"timestamp"`
}

response := APIResponse{
    Timestamp: time.Now().UTC().Format(time.RFC3339),
}
// Will output something like: "2024-03-27T15:04:05Z"
```

## 4. Converting to User's Timezone (when needed):

```go
// RIGHT WAY - Convert UTC to user's timezone only for display
func GetUserTime(utcTime time.Time, userTimezone string) (time.Time, error) {
    location, err := time.LoadLocation(userTimezone)
    if err != nil {
        return time.Time{}, err
    }
    return utcTime.In(location), nil
}

// Usage:
utcTime := time.Now().UTC()
userTime, err := GetUserTime(utcTime, "Asia/Baku")
if err != nil {
    log.Printf("Invalid timezone: %v", err)
}
```
        
### Common Timezone Formats:

```go
// Common formats with timezone information
formats := map[string]string{
    "UTC ISO8601": time.RFC3339,                // "2024-03-27T15:04:05Z"
    "With Zone":   "2006-01-02 15:04:05 MST",   // "2024-03-27 15:04:05 UTC"
    "With Offset": "2006-01-02 15:04:05 -0700", // "2024-03-27 15:04:05 +0000"
}
```
        
### IMPORTANT TIPS:

Always store UTC in your database   
Always use UTC for internal server operations  
Only convert to specific timezones when displaying to users  
Use RFC3339 for API responses   
Be explicit about timezone in logs   

