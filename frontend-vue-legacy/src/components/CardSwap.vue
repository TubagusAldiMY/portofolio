<template>
  <div
      ref="containerRef"
      class="card-swap-container relative"
      :style="{
      width: typeof width === 'number' ? `${width}px` : width,
      height: typeof height === 'number' ? `${height}px` : height,
      perspective: '1200px' /* Menggunakan inline style untuk perspective agar aman */
    }"
  >
    <div
        v-for="(_, index) in 3"
        :key="index"
        ref="cardRefs"
        class="card-swap-card absolute top-1/2 left-1/2 rounded-2xl shadow-2xl overflow-hidden cursor-pointer will-change-transform"
        :style="{
        width: typeof width === 'number' ? `${width}px` : width,
        height: typeof height === 'number' ? `${height}px` : height
      }"
        @click="handleCardClick(index)"
    >
      <slot :name="`card-${index}`" :index="index" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch, nextTick, computed, useTemplateRef } from 'vue';
import gsap from 'gsap';

const props = defineProps({
  width: { type: [Number, String], default: 300 },
  height: { type: [Number, String], default: 400 },
  cardDistance: { type: Number, default: 45 },
  verticalDistance: { type: Number, default: 35 },
  delay: { type: Number, default: 4000 },
  pauseOnHover: { type: Boolean, default: true },
  skewAmount: { type: Number, default: 2 },
  easing: { type: String, default: 'elastic' }
});

const emit = defineEmits(['card-click']);

const containerRef = useTemplateRef('containerRef');
const cardRefs = ref([]);
const order = ref([0, 1, 2]);
const tlRef = ref(null);
const intervalRef = ref(null);

// Helper untuk posisi
const makeSlot = (i, distX, distY, total) => ({
  x: i * distX,
  y: -i * distY,
  z: -i * distX * 1.5,
  zIndex: total - i
});

// Helper set posisi instant
const placeNow = (el, slot, skew) => {
  gsap.set(el, {
    x: slot.x,
    y: slot.y,
    z: slot.z,
    xPercent: -50,
    yPercent: -50,
    skewY: skew,
    transformOrigin: 'center center',
    zIndex: slot.zIndex,
    force3D: true
  });
};

const handleCardClick = (index) => {
  emit('card-click', index);
  swap();
  startAnimation();
};

const config = computed(() => {
  return props.easing === 'elastic'
      ? {
        ease: 'elastic.out(0.6,0.8)',
        durDrop: 1.4,
        durMove: 1.4,
        durReturn: 1.4,
        promoteOverlap: 0.8,
        returnDelay: 0.1
      }
      : {
        ease: 'power2.inOut',
        durDrop: 0.8,
        durMove: 0.8,
        durReturn: 0.8,
        promoteOverlap: 0.45,
        returnDelay: 0.2
      };
});

const initializeCards = () => {
  if (!cardRefs.value.length) return;
  const total = cardRefs.value.length;
  cardRefs.value.forEach((el, i) => {
    if (el) {
      placeNow(el, makeSlot(i, props.cardDistance, props.verticalDistance, total), props.skewAmount);
    }
  });
};

const updateCardPositions = () => {
  if (!cardRefs.value.length) return;
  const total = cardRefs.value.length;
  cardRefs.value.forEach((el, i) => {
    if (el) {
      const slot = makeSlot(i, props.cardDistance, props.verticalDistance, total);
      gsap.set(el, {
        x: slot.x,
        y: slot.y,
        z: slot.z,
        skewY: props.skewAmount
      });
    }
  });
};

const swap = () => {
  if (order.value.length < 2) return;
  const [front, ...rest] = order.value;
  const elFront = cardRefs.value[front];
  if (!elFront) return;

  const tl = gsap.timeline();
  tlRef.value = tl;

  tl.to(elFront, {
    y: '+=350',
    rotation: -5,
    duration: config.value.durDrop,
    ease: config.value.ease
  });

  tl.addLabel('promote', `-=${config.value.durDrop * config.value.promoteOverlap}`);

  rest.forEach((idx, i) => {
    const el = cardRefs.value[idx];
    if (!el) return;
    const slot = makeSlot(i, props.cardDistance, props.verticalDistance, cardRefs.value.length);
    tl.set(el, { zIndex: slot.zIndex }, 'promote');
    tl.to(
        el,
        {
          x: slot.x,
          y: slot.y,
          z: slot.z,
          duration: config.value.durMove,
          ease: config.value.ease
        },
        `promote+=${i * 0.15}`
    );
  });

  const backSlot = makeSlot(cardRefs.value.length - 1, props.cardDistance, props.verticalDistance, cardRefs.value.length);
  tl.addLabel('return', `promote+=${config.value.durMove * config.value.returnDelay}`);

  tl.call(() => { gsap.set(elFront, { zIndex: backSlot.zIndex }); }, undefined, 'return');
  tl.set(elFront, { x: backSlot.x, z: backSlot.z, rotation: 0 }, 'return');

  tl.to(elFront, {
    y: backSlot.y,
    duration: config.value.durReturn,
    ease: config.value.ease
  }, 'return');

  tl.call(() => { order.value = [...rest, front]; });
};

const startAnimation = () => {
  stopAnimation();
  intervalRef.value = window.setInterval(swap, props.delay);
};

const stopAnimation = () => {
  tlRef.value?.kill();
  if (intervalRef.value) clearInterval(intervalRef.value);
};

const setupHoverListeners = () => {
  if (props.pauseOnHover && containerRef.value) {
    containerRef.value.addEventListener('mouseenter', stopAnimation);
    containerRef.value.addEventListener('mouseleave', startAnimation);
  }
};

const removeHoverListeners = () => {
  if (containerRef.value) {
    containerRef.value.removeEventListener('mouseenter', stopAnimation);
    containerRef.value.removeEventListener('mouseleave', startAnimation);
  }
};

watch(() => props.delay, startAnimation);
watch(() => [props.cardDistance, props.verticalDistance, props.skewAmount], updateCardPositions);

onMounted(() => {
  nextTick(() => {
    initializeCards();
    startAnimation();
    setupHoverListeners();
  });
});
onUnmounted(() => {
  stopAnimation();
  removeHoverListeners();
});
</script>

<style scoped>
/* Hanya utilitas standar, tidak ada selector aneh */
.will-change-transform {
  will-change: transform;
}
/* Menghapus class perspective-[1200px] yang menyebabkan error */
/* Style backface-visibility standar */
.card-swap-card {
  backface-visibility: hidden;
  -webkit-backface-visibility: hidden;
}
</style>