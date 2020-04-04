<template>
    <div>
        <span>Twoje doświadczenie: {{ level }} XP</span><br />
        <span>Twój poziom:</span>
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
