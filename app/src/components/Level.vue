<template>
    <div>
        <span>Your level: {{ Math.round(level/100) }} lvl</span><br />
        <span>Your exp: {{ level%100 }} XP</span><br />
        <span>Exp to next level: {{ 100-level%100 }} XP</span>
    </div>
</template>

<script>
  export default {
    name: 'Level',

    data: () => ({
    }),

    props: [ 'level' ],
  }
</script>
