# Implementing Interfaces With Golang

## Learn how to leverage the power of interfaces

[![Yair Fernando](https://miro.medium.com/fit/c/48/48/0*RyyNAe-8KRVc832N.jpg)](https://yairfernando.medium.com/?source=post_page-----51a3b7f527b4--------------------------------)

[Yair Fernando](https://yairfernando.medium.com/?source=post_page-----51a3b7f527b4--------------------------------)

Follow

[Feb 22](https://medium.com/better-programming/implementing-interfaces-with-golang-51a3b7f527b4?source=post_page-----51a3b7f527b4--------------------------------) · 7 min read

![Image for post](https://miro.medium.com/max/30/1*t5KNabwstG5jM4DuB3N7rQ.png?q=20)

![Image for post](https://miro.medium.com/max/1200/1*t5KNabwstG5jM4DuB3N7rQ.png)

Interfaces are tools to define sets of actions and behaviors. They help objects to rely on abstractions and not on concrete implementations of other objects. We can compose different behaviors by grouping multiple interfaces.

In this article we are going to talk about interfaces in Golang and how to use them.

# What Is an Interface?

> An interface is a set of methods that represent common behavior for different data types.

With interfaces, we can organize different groups of methods that apply to different types of objects. By doing this our program can rely on higher abstractions (interfaces) as opposed to concrete implementations, allowing other methods to work with a variety of different objects that implement the same interface.

In the world of OOP, this concept is called the Dependency Inversion Principle. If you want to read more about SOLID principles I recommend you to take a look at [this article.](https://medium.com/swlh/what-are-solid-principles-in-software-development-world-a5ec98637c01)

In Go, it’s considered best practice to build small interfaces and then compose them together to add more functionality to your objects. This way you can keep your code clean and increase reusability.

We can define interfaces considering the different actions that are common between multiple types.

In Go, we can automatically infer that a struct (object) implements an interface when it implements all its methods.

# Defining a Simple Interface in Go

Let’s define an interface in Go and start playing around to discover its power.

type Printer interface {
  Print()
}

This is a really simple interface that defines a method called `Print()`. This method represents the action or behavior that other objects can implement.

Just to be clear, interfaces only define behaviors, they do not define concrete implementations. That’s the job of the object that implements this interface.

Now let’s create two objects that will implement this `Printer` interface:

In the example above we declare two `struct` types — a user and a document.

Then, we use receiver functions to declare the `Print` function on each struct type with their custom implementation.

After this, we can say that both structs are implementing the `Printer` interface.

And by doing this we can write more code that will rely on the abstraction and not in the concrete object, allowing our code to be reused. Let’s say that we want to write a new method that will print the details of these two structs. We can do that using this interface:

func Process(obj Printer) {
    obj.Print()
}

This function takes as an argument any object that implements this interface. So as long as the object responds to the methods defined inside the interface, we can use this function to process the object.

In the main function, we can write the following to print the details of each object:

func main() {
  u := User{name: "John", age: 24, lastName: "Smith"}
  doc := Document{name: "doc.csv", documentType: "csv", date: time.Now()}
  Process(u)
  Process(doc)
}

The output of this code looks like this:

Hi I am John Smith and I am 24years old
Document name: doc.csv, type: csv, date: 2009\-11\-10 23:00:00 +0000 UTC m=+0.000000001

Now, let’s look at a more complex example using interfaces, letting us better understand better how interfaces work in Go.

# Project Description

In this project, we’re going to process client orders. Our program will support `National` and `International` orders and both will depend on the abstraction of an interface to define the expected behavior.

We’ll be using Go modules for this project and I expect you have a basic knowledge of how modules work in Go, but if you don’t it’s not a big deal.

Let’s get started by creating a new folder called `interfaces`.

Inside this folder run the following command to create a module:

go mod init interfaces

This command will generate a new file `go.mod` that includes the name of the module and the Go version. In my case, the go version is `go 1.15`.

Create a new folder called `order`. Inside this folder create the following files:

*   `intenationalOrder.go`
*   `nationalOrder.go`
*   `order.go`
*   `helpers.go`

In the root folder create a new `main.go` file. The folder structure will look like this:

![Image for post](https://miro.medium.com/max/30/1*ws61oUgWmfc9i_MrCr-ruA.png?q=20)

![Image for post](https://miro.medium.com/max/664/1*ws61oUgWmfc9i_MrCr-ruA.png)

Now, let’s take a look at the implementation of the `main.go` file. The definition of the main function is going to be pretty simple since we import the orders package and only call the `New` function from it. In turn, that package will have all the logic for this example:

As you can tell, this file is quite simple, we’re just importing the order package and calling the `New` function on it. We haven’t created the order package but we’ll do so shortly.

For the order package, we’re going to create different struct types and interfaces. Let’s see what this file looks like.

Let’s walk through this file explaining each function, interface, and struct object defined.

First, we have a `New` function. As you already know, we’re using uppercase for the function name because we want to export and make it available to other packages. The purpose of this first function is to create a new instance of a national order and another one for an international one. Then, we pass these two instances to the `ProcessOrder` function inside a slice of type `Operations`. We will discuss this `Operation` type in more detail shortly.

The following struct types represent the various objects that we need to create an order: `Product`, `ProductDetail`, `Summary`, `ShippingAddress`, `Client`, and `Order`.

The `Order` struct type will have a summary, shipping address, and client properties. It also has a products array of type `ProductDetail`.

We also declared three small interfaces: `Processer`, `Printer`, and `Notifier`. Each of these interfaces has one function that defines the behavior other objects have to adopt to implement them.

We have another interface called `Operations`. We’re composing different interfaces to create this one, which is pretty handy because it allows our program to compose objects and make the code more reusable.

Finally, for this file, we have a `ProcessOrder` function which receives an array of orders. Here’s the interesting part. As opposed to receiving an array of a specific objects, this function receives an abstraction of these objects. So, as long as the objects that we pass inside the array implement the `Operations` interface, this function will work correctly. This is where interfaces are really useful, because they allow our program to depend on abstractions and not on concrete implementations.

Now let’s implement the `internationalOrder.go` file:

This file is the first concrete implementation of the `Operations` interface. First, we created a new struct type called `InternationalOrder`, using the `Order` struct to define its properties and objects. Then we have an initializer function called `NewInternationalOrder` that will set some products for this order, the client information, and the shipping address.

We’ll be using a helper function to initialize a new `ProductDetail`, `Client`, and `ShippingAddress` — don’t worry, we’ll implement this soon.

In the rest of the file, we declare the concrete implementation for the `FillOrderSummary`, `Notify`, and `PrintOrderDetails` functions. With this we can say that the international order struct type implements the `Operations` interface because it has definitions for all its methods. Pretty cool!

Let’s take a look at the implementation of the `nationalOrder.go` file:

This file represents the second concrete implementation of the `Operations` interface. Here, we have a `NationalOrder` struct type that also uses the `Order` struct type.

We also declare an initializer function that will set some products, the client information, as well as the shipping address for this particular national order.

Then as we did in the previous file, we have definitions for all methods we need to implement the interface. With this, the national order struct is also implementing the `Operations` interface because it responds to all its methods.

With these two concrete implementations in place, we can pass any of these instances to any methods that rely on the `Operations` interface.

To complete this example we just need to implement the helper functions inside the `helpers.go` file:

As mentioned, this file includes some helper functions to create products, set the client, the shipping address, and calculate the subtotal for the order.

We are ready to run this program!

If we go to the root folder of the project and run the program with `go run main.go`. We should see this expected output:

![Image for post](https://miro.medium.com/max/30/1*JCWlhFBhBTw9SSkOQTV_WA.png?q=20)

![Image for post](https://miro.medium.com/max/1082/1*JCWlhFBhBTw9SSkOQTV_WA.png)

Now that we have learned how to use interfaces in Go, we can create more functions that rely on other interfaces.

For instance, we could create new functions that will receive any objects that implement the `OrderProcesser` or the `OrderNotifier` interface. This is when the idea of defining small interfaces comes in handy.

This would be useful when we have objects that implement smaller interfaces but which are also part of bigger ones. By creating small interfaces, we follow the interface segregation principle, ensuring our program works with different objects that implement different interfaces.

# Conclusion

Interfaces are a great way of creating abstractions that define behavior. Now we can start building programs that use interfaces to abstract behavior and share common actions across different objects. Interfaces also give the program more flexibility and allow more code reuse.

# Resources

If you are interested in learning more about Go, the following articles may be helpful.
