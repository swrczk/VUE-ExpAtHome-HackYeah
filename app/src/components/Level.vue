<template>
    <div class="v-card v-sheet theme--light pa-4" v-if="level>=0">
        <v-layout row wrap class="mb-3">
            <v-flex xs12>
                <v-icon class="mb-1 ml-2" color="orange"> star</v-icon>
                {{ level }} XP
            </v-flex>
        </v-layout>

        <template class="pa-2">
            <v-progress-linear :value="level * 100 / nextLevel"></v-progress-linear>
        </template>

        <v-layout  row wrap class="mb-3">
            <v-flex absolut right xs12>
                <p class="text-lg-right orange--text pa-2">
                    {{ nextLevel - level }} XP to next level
                </p>
            </v-flex>
        </v-layout>
    </div>
    <div class="v-card v-sheet theme--light pa-4" v-else>
        <v-layout row wrap  class="mb-3 text-center">
            <v-flex xs12 class="pa-3">
                Your are so anonymous...<br> If it's your first visit <br> <P class="orange--text">Sign in</p> for more fun.
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    export default {
        name: 'Level',

        data: () => ({}),

        props: ['level'],

        computed: {
            nextLevel() { return (this.level - (this.level % 100)) + 100 },
        },
    }
</script>
