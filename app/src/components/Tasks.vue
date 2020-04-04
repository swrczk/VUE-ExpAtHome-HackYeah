<template>
    <v-layout row wrap>
        <v-flex v-for="(task, t) in tasks" :key="t">
            <v-card class="mb-9">
                <div class="d-flex flex-no-wrap " v-on:click="select">

                    <v-col cols="auto">
                        <v-list-item-avatar
                                tile
                                size="80"
                                color="pink"
                        ></v-list-item-avatar>
                    </v-col>
                    <div>
                        <v-card-title>{{ task.name }}</v-card-title>

                        <v-card-subtitle>{{ task.category }}</v-card-subtitle>
                        <v-card-text>{{ task.info }}</v-card-text>
                    </div>

                </div>

                <v-btn absolute top right fab color="pink">
                    <v-icon> bookmark_border</v-icon>
                </v-btn>

                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn text color="green" @click.prevent="saveTask()">
                        <v-icon> done</v-icon>
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-flex>
    </v-layout>
</template>

<script>
    export default {

        name: 'Tasks',

        data: () => ({}),
        methods: {
            saveTask: function () {
                //bus.$emit('saveTask', true) // emit the event to the bus
                console.log("save")
            },
            select: function () {
                //bus.$emit('saveTask', true) // emit the event to the bus
                console.log("select")
            },
        },
        props: ['tasks'],
    }
</script>
