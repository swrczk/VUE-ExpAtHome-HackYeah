<template>
    <div class="v-card v-sheet theme--light pa-4" v-if="level">
        <v-layout row wrap class="mb-3">
            <v-flex xs12>
                <v-icon class="mb-1 ml-2" color="orange"> star</v-icon>
                {{ level }} lvl
            </v-flex>
        </v-layout>

        <template class="pa-2">
            <v-progress-linear value="40"></v-progress-linear>
        </template>

        <v-layout  row wrap class="mb-3">
            <v-flex absolut right xs12>
                <p class="text-lg-right orange--text pa-2">
                    Needs {{ 100-level%100 }} XP
                </p>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
    export default {
        name: 'Level',

        data: () => ({}),

        props: ['level'],
    }
</script>
