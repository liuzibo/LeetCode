package com.zeber.P0;

import java.io.File;

public class Solution {
    public void hello() {
        System.out.println("Hello World");
    }

    public static void main(String[] args) {
        File dir = new File(".");
        File[] files = dir.listFiles();
        for (File file : files) {
            System.out.println(file.getName());
        }
    }
}
