<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
      clipped-left
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
        />

        #ExpAtHome
      </div>

      <v-spacer></v-spacer>
      <v-toolbar light dense>
        <v-app-bar-nav-icon @click="nav = !nav"></v-app-bar-nav-icon>

        <v-toolbar-title>Title</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-btn icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>

        <v-btn icon>
          <v-icon>mdi-heart</v-icon>
        </v-btn>

        <v-btn icon>
          <v-icon>mdi-dots-vertical</v-icon>
        </v-btn>
      </v-toolbar>

    </v-app-bar>
    <v-navigation-drawer
            :permanent="$vuetify.breakpoint.mdAndUp"
            v-model="nav"
            app
            clipped
    >
      <User :user="user" />
      <Level :level="user.level" />
      <v-divider />
      <v-list>
        <v-list-item
                v-for="item in items"
                :key="item.title"
                link
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <div class="pa-2">
          <v-btn block>Logout</v-btn>
        </div>
      </template>
    </v-navigation-drawer>


    <v-content>
      <v-container>
          <Tasks :tasks="tasks" />
        </v-container>
    </v-content>


    <v-footer>
      <v-spacer></v-spacer>
      <div>HackYeah 2020</div>
    </v-footer>

  </v-app>

</template>

<script>
import User from './components/User';
import Level from './components/Level';
import Tasks from './components/Tasks';

import user from './json/user.json';
import tasks from './json/tasks.json';

export default {
  name: 'App',

  components: {
    User,
    Level,
    Tasks
  },
  theme: {
    dark: true,
  },
  data: () => ({
    items: [
        { title: 'Dashboard', icon: 'dashboard' },
        { title: 'Account', icon: 'account_box' },
        { title: 'Admin', icon: 'gavel' },
    ],
    nav: null,
    user: user,
    tasks: tasks,
  }),
};
</script>
