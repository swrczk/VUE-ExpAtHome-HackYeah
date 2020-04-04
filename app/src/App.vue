<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
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

        NAME
      </div>

      <v-spacer></v-spacer>
      <v-toolbar light dense>
        <v-app-bar-nav-icon></v-app-bar-nav-icon>

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
            permanent
    >
      <!--<v-list>-->
        <!--<v-list-item-->
                <!--v-for="item in items"-->
                <!--:key="item.title"-->
                <!--link-->
        <!--&gt;-->
          <!--<v-list-item-icon>-->
            <!--<v-icon>{{ item.icon }}</v-icon>-->
          <!--</v-list-item-icon>-->

          <!--<v-list-item-content>-->
            <!--<v-list-item-title>{{ item.title }}</v-list-item-title>-->
          <!--</v-list-item-content>-->
        <!--</v-list-item>-->
      <!--</v-list>-->

      <!--<template v-slot:append>-->
        <!--<div class="pa-2">-->
          <!--<v-btn block>Logout</v-btn>-->
        <!--</div>-->
      <!--</template>-->
    </v-navigation-drawer>


      <HelloWorld/>


    <v-footer>
      <v-spacer></v-spacer>
      <div>HackYeah 2020</div>
    </v-footer>

  </v-app>

</template>

<script>
import HelloWorld from './components/HelloWorld';

export default {
  name: 'App',

  components: {
    HelloWorld,
  },

  data: () => ({
    //
  }),
};
</script>
