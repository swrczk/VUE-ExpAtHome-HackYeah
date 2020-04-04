<template>
    <v-row>
        <v-col cols="12" sm="12" md="12" lg="12" >
            <v-text-field 
                outlined
                label="Search tasks" 
                prepend-inner-icon="mdi-magnify"
                v-model="searchBar" 
                hide-details="auto" 
            ></v-text-field>
        </v-col>
        <v-col cols="12" sm="12" md="6" lg="4" v-for="(task, t) in get" :key="t">
            <v-card class="mt-9  pa-2" :color="task.done==0? grey : white">
                <div class="d-flex flex-no-wrap "  @click="task.overlay = !task.overlay">

                    <v-col centre cols="auto">

                        <v-card-text tile size="80" class="orange--text font-weight-bold custom-selector">
                        +{{task.exp}}
                        </v-card-text>
                    </v-col>
                    <div>
                        <v-card-title>{{ task.name }}</v-card-title>
                        <v-card-subtitle bottom>{{ task.username }}</v-card-subtitle>
                        <v-card-subtitle>{{ task.category }}</v-card-subtitle>

                    </div>

                </div>

                <v-btn absolute top right fab color="orange" v-on:click="$emit('savetask', task.id)" v-if=" !task.onlist">
                    <v-icon> bookmark_border</v-icon>
                </v-btn>

                <v-dialog v-model="dialog" persistent max-width="600px"  v-if="task.done==0">
                    <template v-slot:activator="{ on }">
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn text color="green"
                                   v-on="on">
                                <v-icon> done</v-icon>
                            </v-btn>
                        </v-card-actions>
                    </template>

                    <v-card>
                        <v-card-title>
                            <span class="headline">How do you like it?</span>
                        </v-card-title>
                        <v-card-text>
                            <v-container>
                                <v-col cols="12" >
                                <v-row  justify="center">

                                        <v-rating
                                                v-model="rating"
                                                background-color="orange lighten-3"
                                                color="orange"
                                                large
                                                required
                                        ></v-rating>
                                    </v-row>
                                </v-col>
                                    <v-row>
                                    <v-col cols="12" >
                                        <v-text-field  v-model="comment" label="Comment*" hint="Write a few words..."  ></v-text-field>

                                        <v-file-input multiple label="Show us your hard work! :)"></v-file-input>
                                    </v-col>

                                </v-row>
                            </v-container>
                            <small>*indicates required field</small>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="blue darken-1" text @click="dialog = false">Close</v-btn>
                            <v-btn color="blue darken-1" text @click="endTask(task)">Save</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>



                <v-overlay
                        :absolute="absolute"
                        :value="task.overlay"
                >
                    <v-card-actions class="justify-center">
                        <v-card-text v-scroll:#scroll-target="onScroll">{{ task.info }}</v-card-text>
                        <v-btn
                                color="success"
                                @click="task.overlay = false"
                        >
                            Hide
                        </v-btn>
                    </v-card-actions>
                </v-overlay>
            </v-card>

        </v-col>
    </v-row>
</template>

<style >
    .custom-selector{
    font-size: 3em;
    }
</style>

<script>
    export default {

        name: 'Tasks',

        data: () => ({
            comment:"",
            absolute: true,
            rating:4,
            dialog: false,
            searchBar: ""
        }),
        methods: {
            onScroll (e) {
                this.offsetTop = e.target.scrollTop
            },
            endTask: function (task) {
                this.dialog = false;console.log(this.rating)
                this.$emit('endtask', {id:task.id,score:this.rating,comment: this.comment})
            }
        },
        props: ['tasks'],
        computed: {
            get: function () {
                return this.tasks.filter(tasks => {
                    return tasks.name.toLowerCase().includes(this.searchBar.toLowerCase()) ||  tasks.info.toLowerCase().includes(this.searchBar.toLowerCase())
                })
            },
        },
        
    }
</script>
