// Page for problem difficulty
//      @path: /difficulty

<template>
    <div>
        <div class="container">
            <div class="row flex-column justify-content-center align-items-center main-row">
                <mini-calendar :minify="true"/>
                <radio-box :items="this.radioItems" @radioClicked="onRadioClicked"/>
            </div>
            <div class="row main-row">
                <problem-box :problem="this.problem" :titlePrefix="this.radioItems[this.selectedDifficulty]"/>
                <discussion-box @langChanged="onLanguageChanged" :discussion="this.discussion"/>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import td from 'typedef';
import axios from 'axios';

import MiniCalendar from '../components/MiniCalendar.vue';
import RadioBox from '../components/RadioBox.vue';
import ProblemBox from '../components/ProblemBox.vue';
import DiscussionBox from '../components/DiscussionBox.vue';

export default defineComponent({
    name: "Difficulty",
    components: {
        MiniCalendar,
        RadioBox,
        ProblemBox,
        DiscussionBox,
    },
    data() {
        return {
            radioItems: ["easy", "medium", "hard"],
            selectedDifficulty: 0,
            problem: {} as td.ProblemStruct,
            discussion: {} as td.DiscussionStruct,
        };
    },
    methods: {
        /**
         * GET discussion for selected language
         * @params lang: string
         */
        onLanguageChanged(lang: string): void {
            axios.get("/api/difficulty/discussion?lang=" + lang)
            .then((resp) => {
                const data = resp.data as td.ResponseStruct;
                if(typeof data.error === "string") {
                    throw new Error(data.error);
                } else {
                    this.discussion = data.content as td.DiscussionStruct;
                }
            }).catch((err) => {
                alert(err);
            });
        },
        /**
         * GET problem for selected difficulty
         * @params idx: number
         */
        onRadioClicked(idx: number): void {
            axios.get("/api/difficulty/problem?level=" + idx)
            .then((resp) => {
                const data = resp.data as td.ResponseStruct;
                if(typeof data.error === "string") {
                    throw new Error(data.error);
                } else {
                    this.problem = data.content as td.ProblemStruct;
                }
            }).catch((err) => {
                alert(err);
            });
            this.selectedDifficulty = idx;
        },
    },
    created(): void {
        // Init page with first radio item clicked
        this.onRadioClicked(0);
    }
});
</script>

<style scoped>
.main-row {
    min-height: 192px;
}
</style>
