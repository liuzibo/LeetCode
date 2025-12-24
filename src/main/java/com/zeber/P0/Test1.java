package com.zeber.P0;

class Animal {
    public String name = "Animal";
    public void run(){
        System.out.println("Animal run");
    }
}
class Dog extends Animal {
    public String name = "Dog";
    @Override
    public void run(){
        System.out.println("Dog run");
    }
}

class Cat extends Animal {
    public String name = "Cat";
    @Override
    public void run(){
        System.out.println("Cat run");
    }
}

public class Test1 {
    public static void main(String[] args) {
        Animal animal = new Dog();
        System.out.println(animal.name);
        animal.run();
    }
}