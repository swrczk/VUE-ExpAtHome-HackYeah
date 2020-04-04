<template>
<div>
    <v-card v-for="(task, t) in tasks" :key="t">
        <v-card-title>{{ task.name }}</v-card-title>
        <v-card-text>{{ task.info }}</v-card-text>
    </v-card>
</div>
</template>

<script>
  export default {
    name: 'Tasks',

    data: () => ({
    }),

    props: [ 'tasks' ],
  }
</script>
