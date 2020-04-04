<template>
    <v-card>
        <v-card-title vi-if="user">Hi {{ user.name }}!</v-card-title>
        <v-card-title vi-else>Hello mysterious guy!</v-card-title>
    </v-card>
</template>

<script>
  export default {
    name: 'User',

    data: () => ({
    }),

    props: [ 'user' ],
  }
</script>
