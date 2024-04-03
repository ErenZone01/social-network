<script setup lang="ts">
import { RouterView } from "vue-router";
</script>

<template>
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />

    <!-- Favicon -->
    <link href="./src/assets/images/favicon.png" rel="icon" type="image/png" />

    <!-- title and description-->
    <title>Jokkoo</title>
    <meta
      name="description"
      content="Socialite - Social sharing network HTML Template"
    />

    <!-- css files -->
    <link rel="stylesheet" href="./src/assets/css/tailwind.css" />
    <link rel="stylesheet" href="./src/assets/css/style.css" />

    <!-- google font -->
    <link
      href="https://fonts.googleapis.com/css2?family=Inter:wght@200;300;400;500;600;700;800&display=swap"
      rel="stylesheet"
    />
  </head>
  <RouterView />
</template>

<style scoped></style>
