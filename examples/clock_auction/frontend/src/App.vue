<template>
  <v-app>

    <BlockchainViewer v-bind:primary-drawer="primaryDrawer"/>

    <v-app-bar
            :clipped-left="true"
            app
    >

      <v-btn text icon color="purple" @click.stop="primaryDrawer.model = !primaryDrawer.model">
        <v-icon>fa-cubes</v-icon>
      </v-btn>

      <v-toolbar-title class="headline text-uppercase">
        <span>Fabric Private Chaincode</span>
        <span class="font-weight-light"> - Clock Auction Demo</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
              text
              href="https://github.com/hyperledger-labs/fabric-private-chaincode"
              target="_blank"
      >
        <span class="mr-2">Latest Release</span>
      </v-btn>
    </v-app-bar>

    <v-content>
      <MainMenu v-bind:dashboard="dashboard"/>
    </v-content>

  </v-app>
</template>

<script>
  import MainMenu from './components/MainMenu';
  import BlockchainViewer from './components/BlockchainViewer';

  export default {
    name: 'App',
    components: {
      MainMenu,
      BlockchainViewer,
    },
    data: () => ({
      dashboard: {
        tabs: [
          {text: "Elbonia Communication Commission"},
          {text: "A-Telecom"},
          {text: "B-Net"},
          {text: "C-Mobil"},
        ]
      },
      primaryDrawer: {
        model: null,
        blocks: [
          {id: 1, title: "block1"},
          {id: 2, title: "block2"},
          {id: 3, title: "block3"},
          {id: 4, title: "block4"},
          {id: 5, title: "block5"},
          {id: 6, title: "block6"},
        ]
      }
    }),
  };
</script>
