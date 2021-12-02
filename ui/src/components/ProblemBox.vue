// Problem Box component
//      @props:
//          titlePrefix: string?
//              Prefix of the header-2
//          problem: ProblemStruct
//              Object containing title, link and th description of the problem
//      @emit:
//          none

<template>
    <div class="col no-pad">
        <h2 class="fira-sans font-bold">{{this.refinedPrefix}}Problem of the Day</h2>
        <article v-if="this.problemReady">
            <a :href="this.problem.link" target="_blank"><h3>{{this.problem.title}}</h3></a>
            <div id="fetched-problem" v-html="this.problem.description"></div>
        </article>
        <article v-else>
            <spinner-box/>
        </article>
    </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import td from 'typedef';
import SpinnerBox from './SpinnerBox.vue'

export default defineComponent({
    name: "ProblemBox",
    props: {
        titlePrefix: String,
        problem: Object as PropType<td.ProblemStruct>,
    },
    components:{
        SpinnerBox,
    },
    computed: {
        /**
         * Check if title prefix is undefined
         * @return string
         */
        refinedPrefix(): string {
            return this.titlePrefix ? this.titlePrefix.toUpperCase() + ' ' : "";
        },
        problemReady(): boolean {
            if(this.problem instanceof Object) {
                return Object.keys(this.problem).length > 0;
            }
            return false;
        }
    },
    updated() {
        // Redefine size of <img> tag to fit in
        document.querySelectorAll("img").forEach((el) => {
            el.style.width = "500px";
            el.style.height = "auto";
        });
    }
});
</script>

<style scoped>
h2 {
    font-size: 28px;
    margin: 0px 15px 8px 15px;
    border-bottom: 1px solid black;
}

h3 {
    font-size: 24px;
}

article {
    padding: 0px 15px;
    border-right: 1px solid black;
}
</style>
