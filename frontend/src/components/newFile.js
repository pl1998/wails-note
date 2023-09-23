import { computed, defineProps } from "vue";
import { store } from './../store/index';

defineProps({
themes: ["default", "dark", "light"],
themePath: computed(() => {
return `./theme-${store.state.theme}.css`;
}),
theme: store.state.theme,
});
