import { createApp, h } from 'vue'
import { createInertiaApp } from '@inertiajs/vue3'

createInertiaApp({

  resolve: name => {
    const pages = import.meta.glob("./Pages/**/*.vue", { eager: true });
    const page = pages[`./pages/${name}.vue`] as any;
    page.default.layout = page.default.layout;
    return page;
  },

  setup({ el, App, props, plugin }) {
    createApp({ render: () => h(App, props) })
      .use(plugin)
      .mount(el)
  },

})