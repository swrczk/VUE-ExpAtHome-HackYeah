<template>
    <v-card v-if="user.name">
        <v-card-title>Hi {{ user.name }}!</v-card-title>
    </v-card>
    <v-card v-else>
        <v-card-title >Hello mysterious guy!</v-card-title>
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
