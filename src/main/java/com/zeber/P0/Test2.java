package com.zeber.P0;

public class Test2 {

    public static void main(String[] args) {
        Animal animal = new Cat();
        if (animal instanceof Dog){
            Dog dog = (Dog) animal;
            System.out.println(dog.name);
        }else if (animal instanceof Cat c){
            System.out.println(c.name);
        }
    }
}


