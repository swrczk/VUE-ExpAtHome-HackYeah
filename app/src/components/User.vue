<template>
    <v-card>
        <v-card-title>{{ user.name }}</v-card-title>
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
