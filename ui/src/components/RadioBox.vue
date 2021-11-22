// Radio component
//      @props: string[]
//          Labels for radio items
//      @emit:
//          event name: radioClicked
//          parameter:
//              int: Index of clicked radio item

<template>
    <ul class="d-inline-block">
        <li v-for="(item, idx) in radioItems" v-bind:key="idx" @click="onClick(idx)" class="d-inline-block">
            <div class="open-sans font-bold text-white selected" v-if="selected == idx">{{item}}</div>
            <div class="open-sans font-bold text-white" v-else>{{item}}</div>
        </li>
    </ul>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';

export default defineComponent({
    name: "RadioBox",
    props: {
        items: Array as PropType<string[]>,
    },
    data() {
        return {
            selected: 0,
        };
    },
    methods: {
        /*
         * Alert to parent when radio item is clicked
         * @params idx: number
         */
        onClick(idx: number): void {
            this.selected = idx;
            this.$emit("radioClicked", idx);
        }
    },
    computed: {
        /*
         * Convert radio items to upper case letter
         * @return string[]
         */
        radioItems(): string[] {
            if(this.items instanceof Array) {
                return this.items.map((val) => val.toUpperCase());
            }
            return [];
        }
    }
});
</script>

<style scoped>
ul {
    list-style: none;
    width: auto;
}

li>div {
    background-color: #A9A9A9;
    padding: 15px 20px;
    font-size: 20px;
    margin-left: 2px;
    cursor: pointer;
}

.selected {
    background-color: #285CBF;
}
</style>
