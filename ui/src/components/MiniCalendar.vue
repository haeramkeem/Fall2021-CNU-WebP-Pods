<template>
    <div class="container-auto no-pad" id="mini-calendar">
        <div class="row align-items-center justify-content-center">
            <div v-if="this.minify" class="col-auto open-sans font-bold" id="mini-scale">{{strWeekDay}}, {{strFullDate}}</div>
            <div v-else class="col-auto no-pad" id="full-scale">
                <button class="arrow-btn" id="yesterday-btn"></button>
                <div class="text-center" id="day-date">
                    <div class="open-sans font-reg" id="week-day">{{strWeekDay}}</div>
                    <div class="fira-sans font-bold" id="full-date">{{strFullDate}}</div>
                </div>
                <button class="arrow-btn" id="tomorrow-btn"></button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { defineComponent } from 'vue';

const date = new Date();

export default defineComponent({
    name: "MiniCalendar",
    props: {
        minify: Boolean,
    },
    data() {
        return {
            year: date.getFullYear() - 2000,
            month: date.getMonth() + 1,
            day: date.getDate(),
            weekDay: date.getDay(),
        };
    },
    computed: {
        strWeekDay() {
            const week = [
                "SUNDAY",
                "MONDAY",
                "TUESDAY",
                "WEDNESDAY",
                "THURSDAY",
                "FRIDAY",
                "SATURDAY",
            ];
            return week[this.weekDay];
        },
        strFullDate() {
            return `${this.year}${this.month}${this.day}`;
        }
    }
});
</script>

<style>
#mini-calendar { width: 400px; }
#full-scale>* {
    display: inline-block;
}

#mini-scale { font-size: 24px; }

#week-day {
    color: grey;
}

#full-date {
    font-size: 60px;
}

.arrow-btn {
    width: 30px;
    height: 44px;
    border: none;
    outline: none;
    background-color: white;
    background-repeat: no-repeat;
    background-size: contain;
}

#yesterday-btn {
    background-image: url("../assets/img/left-arrow.svg");
}

#tomorrow-btn {
    background-image: url("../assets/img/right-arrow.svg");
}

#day-date {
    width: 300px;
}

.text-center * {
    text-align: center;
}
</style>
