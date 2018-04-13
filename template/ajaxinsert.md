# How To Submit a Form with jQuery and AJAX

HTML forms make up a large part of the web. They are the primary method for retrieving input from users.  Typically, you fill out the form, click the submit button, and be redirected to a thank you page.  For web applications, this may not be ideal - you may not want the user to leave the page. In this article, I'll show you how to use jQuery's AJAX function to submit a form asynchronously to the server and avoid a redirect.

## A Simple Contact Form
Here's an example of a simple contact form (fyi: submitting it won't do anything yet). We'll be submitting this form via AJAX with jQuery.

Your Name: 
Your Email: 
Your Message:

 
Here's the code for the form:

```
<form id="ContactForm">
    Your Name: <input type="text" name="name" value="" /><br /> 
    Your Email: <input type="text" name="email" value="" /><br /> 
    Your Message:<br /> <textarea style="width: 200px; height: 100px;" name="message"></textarea> 
    <br /><br /> 
    <input type="submit" name="submit" value="Submit" /><br />
    <div class="form_result"> </div>
</form>
```

Notice that I gave my form tag an ID of "ContactForm". We'll need this later so that jQuery can find and retrieve all the data contained in the form. Additionally, I've included a div element to which we will dynamically update with the response from the server.  The div has a class of 'form_result' that we can use to tell jQuery where the response from the server should be output to.  The reason I'm using a class identifier here instead of giving the div an id attribute is because on very large forms, you may want the result of the form to be shown at the very top of the form as well as the bottom.  With jQuery, we can select and update multiple elements with the same class.  However, we can only update one element with the same id.  

## jQuery's AJAX Function
The syntax for jQuery's ajax function looks like this:

jQuery.ajax(url, [settings]);
A complete reference to this function can be found by clicking on this link:jQuery's AJAX function. There are numerous settings available to use with this function, but we're only going to be using a handful. We'll discuss each setting briefly here: 

url: set this to the URL the form will be submitted to (you'll probably want to put in the URL to your form handler). We'll be setting ours to jQuery-ajax-demo.php for demonstration purposes.
type: this is the method type for the form; set it to POST or GET. In our case, we'll be using POST.
data: this is the data you'll be submitting; it's a standard query string that you typically see after a domain name in the browser's address bar. You'll see that for the most part, jQuery will create the query string for us.
success: you can set this to a single callback function, or as of jQuery 1.5 - an array of functions. We'll be setting ours to a single callback function that accepts the response from the server (first parameter of the callback function).
Let's see what a call to jQuery's AJAX function would look like for our demo:

```
$.ajax({type:'POST', url: 'jQuery-ajax-demo.php', data:$('#ContactForm').serialize(), success: function(response) {
    $('#ContactForm').find('.form_result').html(response);
}});
```

As you can see, we are setting the 'type', 'url', 'data', and 'success' settings of jQuery's AJAX function. In the 'data' setting, we're using jQuery's serialize() function by supplying the id of the form as the context.
jQuery's serialize() function looks at every field in the supplied context and creates a string of key-value pairs. The 'name' attribute of each of the fields in the form are used as the keys while the actual data contained in the fields become the values. So for example, in the form we created above, our query string (the result of the serialize function) might look like this:
```
name=Jason&email=me%40example.com&message=hey+what's+up%3F
```

That's if I submitted the form with Name set to "Jason", Email set to "me@example.com", and Message set to "hey what's up?". The serialize() method also encodes the data so it can be easily transmitted to the server. 
We now have everything we need to submit the form to the server, but we do not have a form handler yet that will intercept the data and do something with it.  

## A Simple Form Handler in PHP
Doing anything too fancy is outside the scope of this article, so we're just going to create a simple form handler that intercepts our data and outputs the data as the response.  Take a look at this PHP code:
```
if(isset($_POST['name'])) {
    $name = $_POST['name'];
    $email = $_POST['email'];
    $msg = $_POST['message'];
	
    ?>
    Your Name Is: <?php echo $name; ?><br />
    Your Email Is: <?php echo $email; ?><br />
    Your Message Is: <?php echo $msg; ?><br />
    <?php
    die();
}
```

Here we are first checking to see if the form was submitted. The result of isset($_POST['name']) will be true if the key 'name' is defined in PHP's $_POST array, and false otherwise. It will be since 'name' is one of the fields on the form that is serialized and submitted via the AJAX request. Next, we retrieve and then output the data from each of the fields: name, email and message. I told you it was simple! In reality though, you'd probably want to validate the data. Although validating data from a form submission is outside the scope of this article, it does give me an idea for an article I could write next!

## Putting It All Together
We have everything we need now, we just have to put it all together! The first thing we will do is create a JavaScript function we can call that runs our jQuery AJAX code, here's what that looks like:

```
function submitForm() {
    $.ajax({type:'POST', url: 'jQuery-ajax-demo.php', data:$('#ContactForm').serialize(), success: function(response) {
        $('#ContactForm').find('.form_result').html(response);
    }});

    return false;
}
```
You can see we return false from the function. This is very important! When we return false from this function, it will stop the browser from submitting the form and reloading the page or redirecting us. The last we have to do is update the code for our form to call our submitForm() function:

```
<form id="ContactForm" onsubmit="return submitForm();">
    Your Name: <input type="text" name="name" value="" /><br /> 
    Your Email: <input type="text" name="email" value="" /><br /> 
    Your Message:<br /> <textarea style="width: 200px; height: 100px;" name="message"></textarea> 
    <br /><br /> 
    <input type="submit" name="submit" value="Submit" /><br />
    <div class="form_result"> </div>
</form>
```
