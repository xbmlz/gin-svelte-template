<script lang="ts">
    import 'uno.css'
    import '@unocss/reset/tailwind.css'

    import Footer from '$lib/components/layout/Footer.svelte';
</script>

<div class="bg-white dark:bg-slate-950">
    <div class="overflow-hidden lg:flex">

        <!-- Sidebar start -->
        <aside class="fixed inset-0 z-30 flex-none h-full w-64 lg:h-auto border-e border-gray-200 dark:border-gray-600 lg:overflow-y-visible lg:pt-16 lg:block hidden">Sidebar</aside>
        <!-- Sidebar end -->

        <div class="fixed inset-0 z-20 bg-gray-900/50 dark:bg-gray-900/60" role="presentation" hidden=""></div>

        <!-- Body start -->
        <div class="relative h-full w-full overflow-y-auto lg:ml-64 pt-[70px]">
            <slot/>

            <Footer/>
        </div>
        <!-- Body end -->
    </div>
</div>

<style>
    :root {
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
        Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    }
</style>