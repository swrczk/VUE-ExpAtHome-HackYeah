<template>
<div>

    <v-col cols="12" sm="12" md="6" lg="4" v-for="(achievement, t) in achievements" :key="t">
        <v-row>

            <v-avatar color="indigo" v-if="achievement.owned">
                <v-icon dark>achievement.icon</v-icon>
            </v-avatar>
            <v-avatar   color="grey" v-else>
                <v-icon dark>achievement.icon</v-icon>
            </v-avatar>

            <v-card-title>{{achievement.title}}}</v-card-title>

        </v-row>
    </v-col>
    </div>
</template>

<script>
    export default {
        name: 'Achievements',

        data: () => ({
            achievements:[
                {"title": "Welcome!","icon": "accessibility_new", "owned": true},
                {"title": "It's totally new field for you!","icon": "3d_rotation", "owned": false},
                {"title": "IT is easy fo me","icon": "android", "owned": true},
                {"title": "Creative way of traveling","icon": "commute", "owned": false},
                {"title": "Seek, and ye shall find.","icon": "contact_support", "owned": true},
                {"title": "Seek, and ye shall find.","icon": "contact_support", "owned": false},
                {"title": "Mother of nature","icon": "eco", "owned": true},
                {"title": "J'adore!","icon": "favorite", "owned": false},
                {"title": "Such a quick start!","icon": "flight_takeoff", "owned": true},
                {"title": "Shine bright like a star","icon": "grade", "owned": false},
                {"title": "Unlock yourself","icon": "lock_open", "owned": false},
                {"title": "Gramar? Im so famijar","icon": "spellcheck", "owned": true},
                {"title": "Math doesn't lie... you rock!","icon": "timeline", "owned": false},
                {"title": "See everything", "icon": "visibility", "owned": false},
                {"title": "Big bag of knowledge", "icon": "work_outline", "owned": false},
                {"title": "Precious commenter", "icon": "library_books", "owned": false},
                {"title": "You're feed helped a lot of people", "icon": "dynamic_feed", "owned": false},
                {"title": "Cloud of ideas", "icon": "settings_system_daydream", "owned": false},
                {"title": "Nice strike!", "icon": "insert_emoticon", "owned": true},
                {"title": "Best guide", "icon": "security", "owned": false},
                {"title": "Your storyline is so inspiring", "icon": "videogame_asset", "owned": false},
                {"title": "First photo?", "icon": "add_photo_alternate", "owned": false},
                {"title": "Night is not my limit", "icon": "brightness_3", "owned": false},
                {"title": "Good shooter", "icon": "camera_alt", "owned": false},
                {"title": "People love your collections", "icon": "collections", "owned": false},
                {"title": "Thank you for being with us :D", "icon": "filter_vintage", "owned": true},
                {"title": "Transformation time!", "icon": "transform", "owned": false},
                {"title": "Category master", "icon": "category", "owned": false},
                {"title": "Amateur Master Chef", "icon": "local_dining", "owned": false},
                {"title": "Own Director", "icon": "local_movies", "owned": true}
            ]
        }),

        props: [],
    }
</script>
