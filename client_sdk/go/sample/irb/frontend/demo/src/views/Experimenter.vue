<template>
  <v-stepper
      :value="currentProgressStep"
      class="elevation-0"
  >
    <v-stepper-header class="elevation-0">
      <v-stepper-step
          :complete="currentProgressStep > 1"
          step="1"
          editable
      >
        Create proposal
      </v-stepper-step>
      <v-divider></v-divider>

      <v-stepper-step
          :complete="currentProgressStep > 2"
          step="2"
          editable
      >
        Wait for approvals
      </v-stepper-step>
      <v-divider></v-divider>

      <v-stepper-step
          :complete="currentProgressStep > 3"
          step="3"
          editable
      >
        Spawn experiment instance
      </v-stepper-step>
      <v-divider></v-divider>

      <v-stepper-step step="4" editable>
        Provision and execute
      </v-stepper-step>

    </v-stepper-header>

    <v-stepper-items>
      <v-stepper-content step="1">
        <ProposalCreate
            class="mb-2"
            v-on:submit-proposal="onSubmitProposal"
        />
      </v-stepper-content>

      <v-stepper-content step="2">
        <WaitingForApprovals
            class="mb-2"
            :watchProposalWithId="currentProposalId"
            v-on:next="onNextStep"
        />
      </v-stepper-content>

      <v-stepper-content step="3">
        <SpawnExperiment
            class="mb-2"
            :watchProposalWithId="currentProposalId"
            v-on:next="onNextStep"
        />
      </v-stepper-content>

      <v-stepper-content step="4">
        <ExecuteExperiment
            :watchProposalWithId="currentProposalId"
            class="mb-2"
        />
      </v-stepper-content>

    </v-stepper-items>
  </v-stepper>
</template>

<script>
// @ is an alias to /src
import ProposalCreate from '@/components/proposal/Create.vue';
import WaitingForApprovals from '@/components/experiment/Waiting.vue';
import SpawnExperiment from '@/components/experiment/Spawn.vue';
import ExecuteExperiment from '@/components/experiment/Execute.vue';


import {mapGetters} from 'vuex';

export default {
  name: 'Experimenter',
  components: {
    ProposalCreate,
    WaitingForApprovals,
    SpawnExperiment,
    ExecuteExperiment,
  },
  data: () => ({
    currentProposalId: '1'
  }),
  computed: {
    ...mapGetters({
      currentProgressStep: 'experiment/currentProgressStep',
    }),
  },
  methods: {
    onSubmitProposal: function () {
      // TODO get proposal ID
      this.onNextStep();
    },
    onNextStep: function () {
      this.$store.dispatch('experiment/nextStep');
    }
  },
};
</script>
