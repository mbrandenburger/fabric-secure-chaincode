<template>
  <v-app id="inspire">
    <v-content>
      <v-container class="fill-height" fluid>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="8" md="4">
            <v-card class="elevation-12">
              <v-toolbar color="primary" dark flat>
                <v-toolbar-title>Spectrum Auction Login</v-toolbar-title>
                <v-spacer></v-spacer>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-col class="d-flex" cols="12">
                    <v-select
                      v-model="selection"
                      :items="usernames"
                      filled
                      label="Registered Users"
                    ></v-select>
                  </v-col>
                  <v-text-field
                    id="password"
                    label="Password"
                    name="password"
                    prepend-icon="fa-lock"
                    type="password"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <router-link to="/auction_info">
                  <v-btn @click="loginUser" color="primary">Login</v-btn>
                </router-link>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";

export default {
  name: "Login",
  props: {
    source: String
  },
  data: () => ({
    userdata: [],
    usernames: [],
    selection: ""
  }),
  computed: {
    ...mapState({
      auction: state  => state.auction.auction
    })
  },
  mounted() {
    axios
      .get("http://localhost:3000/api/getAllUsers")
      .then(response => {
        this.userdata = response.data;
        this.usernames = this.userdata.map(function(el) {
          return el.userName;
        });
      })
      .catch(error => {
        alert("error " + error);
        this.errored = true;
      });
  },
  methods: {
    loginUser: function() {
      var user = this.userdata.filter(
        element => element.userName == this.selection
      )[0];
      //alert(JSON.stringify(user));
      let currentObj = this;
      axios
        .post("http://localhost:3000/api/login", user)
        .then(function(response) {
          console.log(response);
          currentObj.$store
            .dispatch("auction/UPDATE_CURRENT_USER", user.userName)
            .then(() => {
              console.log("current user updated: " + user.userName);
            });
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>
