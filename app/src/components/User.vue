<template>
    <v-card>
        <v-card-title>Hi {{ user.name }}!</v-card-title>
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
