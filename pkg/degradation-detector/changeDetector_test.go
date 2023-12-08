package degradation_detector

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestChangeDetector(t *testing.T) {
  data := []int{5691, 5855, 5720, 6339, 5829, 5496, 5427, 5586, 5859, 5603, 5868, 5761, 5440, 5590, 5870, 5781, 5632, 6092, 5636, 5849, 5730, 5639, 5678, 5857, 5655, 5486, 5877, 5639, 5668, 5864, 5602, 5855, 6049, 5741, 5794, 5822, 5704, 5707, 6167, 5923, 5765, 5648, 5775, 5578, 5541, 5919, 5498, 5436, 5857, 5508, 5739, 5820, 5662, 5582, 5565, 5708, 5587, 5813, 5618, 5796, 5682, 5778, 5848, 6034, 5847, 5653, 5783, 6006, 5647, 5509, 5869, 5738, 5709, 5762, 5793, 5607, 5620, 5580, 5710, 5641, 5673, 5794, 5937, 5708, 5705, 5747, 5679, 5963, 6240, 5958, 5915, 5737, 6000, 5747, 5529, 5562, 5909, 5713, 5680, 5729, 5656, 5820, 5670, 5884, 5686, 5662, 5848, 5710, 5707, 5821, 5564, 6029, 6045, 5765, 5727, 5653, 5766, 5784, 5893, 5755, 5756, 5836, 5652, 5971, 6000, 5689, 6110, 5953, 6102, 5747, 5872, 5808, 5891, 5839, 5719, 5865, 6114, 5811, 5687, 5834, 5759, 5873, 6114, 6314, 5757, 5849, 5901, 5993, 5226, 5338, 5356, 5299, 5426, 5479, 5687, 5594, 5497, 5735, 19094, 19217, 19264, 18976, 19040, 19348, 19092, 5777, 5810, 5636, 5600, 5681, 5528, 5573, 5494, 5613, 5603, 5509, 5455, 5552, 5773, 5903, 5385}
  indexes := GetChangePointIndexes(data, 1)
  assert.Equal(t, []int{111, 148, 153, 158, 165}, indexes)
}

func TestWithoutDegradations(t *testing.T) {
  data := []int{99, 85, 68, 70, 67, 82, 87, 68, 93, 72, 85, 86, 84, 71, 95, 67, 85, 97, 86, 87, 85, 97, 99, 101, 75, 86, 107, 87, 97, 120, 75, 80, 70, 78, 75, 92, 97, 98, 98, 115, 85, 110, 120, 85, 70, 80, 105, 104, 74, 78, 100, 109, 98, 107, 74, 72, 69, 66, 86, 76, 66, 75, 84, 85, 93, 74, 95, 98, 71, 70, 74, 98, 109, 75, 120, 72, 86, 72, 73, 80, 102, 85, 86, 74, 75, 74, 98, 100, 110, 85, 90, 105, 95, 113, 71, 98, 111, 76, 69, 106, 97, 78, 67, 86, 109, 88, 94, 111, 86, 71, 72, 85, 74, 66, 84, 86, 67, 67, 75, 90, 78, 101, 67, 75, 81, 87, 74, 84, 95, 89, 96, 74, 74, 71, 114, 98, 87, 87, 107, 111, 86, 71, 75, 74, 88, 83, 86, 62, 77, 98, 74, 77, 87, 76, 90, 110, 78, 112, 113, 100, 101, 86, 102, 85, 69, 67, 74, 87, 98, 98, 94, 71, 72, 73, 85, 85, 99, 97, 96, 72, 94, 85, 110, 75, 69, 89, 73, 85, 90, 108, 76, 73, 94, 105, 103, 86, 96, 67, 102, 95, 112, 69, 89, 96, 71, 69, 89, 84, 74, 88, 87, 78, 110, 97, 72, 64, 99, 98, 110, 118, 74, 77, 109, 73, 74, 102, 71, 88, 87, 99, 98, 73, 101, 105, 85, 69, 69, 95, 73, 79, 89, 73, 101, 71, 86, 100, 86, 69, 71, 66, 71, 84, 100, 66, 71, 96, 71, 96, 101, 87, 86, 85, 65, 99, 81, 95, 98, 98, 75, 70, 72, 73, 74, 71, 84, 73, 95, 85, 106, 70, 71, 67, 69, 85, 74, 84, 108, 100, 95, 84, 67, 85, 86, 88, 71, 91, 71, 101, 70, 69, 68, 76, 73, 72, 86, 87, 73, 86, 73, 96, 87, 71, 88, 70, 95, 67, 98, 75, 72, 72, 95, 95, 68, 76, 84, 96, 73, 95, 97, 95, 63, 85, 100, 70, 110, 85, 100, 73, 97, 99, 100, 94, 83, 93, 85, 72, 102, 76, 70, 96, 107, 82, 78, 74, 98, 113, 71, 85, 87, 70, 101, 74, 98, 72, 98, 98, 73, 87, 102, 88, 86, 99, 70, 86, 75, 86, 74, 67, 111, 72, 96, 99, 75, 75, 70, 70, 71, 95, 70, 86, 92, 109, 96, 91, 88, 70, 97, 69, 74, 64, 103, 97, 84, 71, 97, 102, 80, 76, 112, 76, 99, 69, 74, 69, 90, 77, 86, 107, 96, 68, 99, 89, 89, 85, 73, 88, 85, 84, 75, 82, 95, 68, 98, 90, 94, 85, 86, 84, 85, 73, 94, 97, 95, 74, 85, 73, 107, 99, 72, 70, 75, 88, 87, 85, 98, 97, 84, 91, 71, 70, 75, 88, 97, 70, 95, 77, 66, 76, 109, 74, 84, 69, 81, 76, 87, 72, 97, 101, 109, 85, 98, 84, 97, 75, 112, 108, 73, 96, 73, 84, 84, 73, 86, 70, 69, 71, 73, 85, 67, 101, 97, 91, 74, 75, 97, 82, 73, 73, 85, 97, 61, 70, 72, 85, 88, 71, 67, 85, 65, 68, 68, 98, 84, 73, 87, 71, 80, 77, 90, 100, 71, 120, 69, 85, 87, 82, 85, 96, 128, 86, 71, 69, 102, 85, 85, 110, 71, 74, 73, 88, 72, 94, 97, 85, 74, 67, 111, 73, 96, 84, 94, 94, 66, 66, 75, 75, 74, 74, 94, 94, 74, 74, 69, 69, 73, 73, 86, 86, 66, 66, 95, 95, 73, 73, 88, 88, 97, 97, 86, 86, 86, 86, 88, 88, 68, 68, 69, 69, 86, 86, 85, 85, 73, 73, 73, 73, 103, 103, 99, 99, 97, 97, 72, 72, 83, 83, 70, 70, 98, 98, 96, 96, 106, 106, 85, 85, 108, 108, 114, 114, 86, 86, 96, 96, 77, 77, 97, 97, 102, 102, 89, 89, 107, 107, 99, 99, 101, 101, 96, 96, 113, 113, 112, 112, 98, 98, 83, 83, 72, 72, 74, 74, 96, 96, 84, 84, 91, 91, 97, 97, 73, 73, 88, 88, 83, 83, 97, 97, 114, 114, 84, 84, 75, 75, 89, 89, 75, 75, 94, 94, 87, 87, 86, 86, 101, 101, 96, 96, 138, 138, 104, 104, 121, 121, 71, 71, 99, 99, 72, 72, 72, 72, 87, 87, 95, 95, 73, 73, 71, 71, 82, 82, 74, 74, 88, 88, 106, 106, 74, 74, 85, 85, 74, 74, 85, 85, 86, 86, 88, 88, 103, 103, 99, 99, 91, 91, 101, 101, 100, 100, 85, 85, 73, 73, 83, 83, 95, 95, 71, 71, 73, 73, 75, 75, 87, 87, 74, 74, 77, 77, 102, 102, 91, 91, 81, 81, 79, 79, 99, 99, 79, 79, 94, 94, 76, 76, 78, 78, 84, 84, 81, 81, 90, 90, 93, 93, 95, 95, 79, 79, 81, 81, 85, 85, 87, 87, 87, 87, 85, 85, 83, 83, 81, 81, 88, 88, 88, 88, 75, 75, 88, 88, 84, 84, 74, 74, 74, 113, 74, 84, 75, 99, 85, 97, 85, 73, 92, 72, 86, 86, 109, 83, 101, 68, 73, 73, 86, 68, 101, 108, 79, 96, 93, 97, 71, 93, 88, 94, 78, 95, 83, 82, 93, 92, 86, 101, 81, 81, 81, 148, 81, 77, 92, 81, 94, 83, 91, 93, 89, 82, 97, 83, 95, 87, 77, 106, 94, 90, 84, 98, 93, 97, 97, 92}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{DoNotReportImprovement: true})
  assert.Equal(t, 0, len(degradations))
}
func TestWithoutDegradations2(t *testing.T) {
  data := []int{2940, 2633, 2758, 2648, 2884, 2920, 3205, 2936, 2868, 3212, 2324, 2290, 2474, 3000, 2740, 2737, 2413, 2873, 3105, 3049, 2521, 3185, 2950, 2696, 2725, 3188, 3781, 2493, 2241, 2528, 2680, 3126, 3126, 2649, 2649, 2623, 2623, 2820, 2820, 2693, 2693, 2812, 2812, 2527, 2527, 3031, 3031, 2571, 2571, 3066, 3066, 2670, 2670, 2699, 2699, 3106, 2782, 2617}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{})
  assert.Equal(t, 0, len(degradations))
}

func TestWithDegradations(t *testing.T) {
  data := []int{101, 99, 101, 99, 101, 201, 202, 201, 202, 201, 201, 201}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 3})
  assert.Equal(t, 1, len(degradations))
  assert.Equal(t, int64(5), degradations[0].timestamp)
}

func TestWithoutDegradations3(t *testing.T) {
  data := []int{100, 100, 100, 100, 100, 200, 100, 100, 100}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 3})
  assert.Equal(t, 0, len(degradations))
}

func TestWithDegradations2(t *testing.T) {
  data := []int{100, 100, 100, 100, 100, 200, 300, 300, 300}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 3})
  assert.Equal(t, 1, len(degradations))
  assert.Equal(t, int64(5), degradations[0].timestamp)
}

func TestComplexDistributionFromUnitTestWithoutDegradation(t *testing.T) {
  data := []int{310, 332, 331, 324, 350, 346, 280, 362, 284, 350, 312, 379, 376, 368, 326, 308, 335, 293, 336, 376, 318, 286, 307, 311, 363, 273, 351, 361, 362, 368, 350, 379, 286, 290, 296, 287, 321, 287, 315, 297, 331, 277, 369, 344, 282, 283, 272, 387, 282, 271, 351, 311, 286, 315, 384, 283, 325, 288, 298, 370, 360, 320, 289, 389, 367, 355, 439, 295, 276, 345, 297, 359, 386, 319, 397, 351, 347, 346, 305, 387, 348, 295, 282, 360, 337, 399, 329, 305, 400, 317, 296, 287, 288, 351, 338, 309, 258, 273, 296, 273, 321, 321, 261, 274, 275, 255, 253, 330, 329, 275, 276, 304, 263, 278, 316, 302, 330, 351, 318, 307, 281, 287, 274, 381, 320, 347, 280, 291, 307, 353, 354, 284, 388, 352, 369, 410, 348, 266, 279, 256, 286, 287, 352, 341, 293, 282, 279, 362, 334, 344, 322, 402, 277, 289, 270, 273, 357, 356, 290, 317, 260, 272, 286, 332, 281, 337, 267, 276, 279, 299, 306, 325, 444, 388, 320, 292, 377, 306, 297, 371, 524, 301, 452, 324, 360, 296, 332, 378, 265, 355, 310, 301, 399, 328, 315, 462, 372, 378, 322, 326, 335, 272, 303, 328, 291, 273, 326, 299, 324, 307, 328, 267, 378, 328, 275, 324, 309, 304, 284, 301, 294, 291, 298, 330, 350, 348, 313, 313, 299, 340, 347, 289, 301, 339, 337, 375, 295, 311, 366, 341, 325, 422, 297, 348, 359, 271, 284, 276, 267, 275, 285, 254, 386, 317, 353, 264, 265, 286, 299, 294, 277, 347, 267, 342, 274, 318, 367, 273, 335, 269, 301, 361, 424, 340, 316, 280, 304, 311, 463, 312, 435, 362, 328, 321, 270, 330, 407, 277, 319, 426, 319, 366, 466, 370, 363, 352, 372, 264, 274, 287, 324, 333, 363, 267, 292, 294, 307, 321, 307, 295, 270, 332, 308, 271, 291, 276, 270, 331, 265, 278, 295, 320, 315, 290, 294, 364, 350, 343, 324, 341, 355, 374, 273, 376, 305, 304, 352, 351, 375, 321, 334, 320, 367, 367, 386, 265, 344, 281, 328, 350, 369, 299, 321, 371, 300, 342, 369, 312, 285, 348, 356, 291, 280, 285, 360, 283, 278, 286, 353, 359, 295, 446, 319, 394, 346, 352, 284, 269, 375, 281, 310, 383, 388, 343, 387, 283, 342, 293, 356, 301, 273, 348, 343, 398, 275, 346, 411, 383, 351, 264, 278, 284, 369, 281, 278, 303, 266, 299, 295, 348, 262, 261, 331, 277, 267, 380, 260, 266, 352, 279, 357, 295, 271, 352, 279, 318, 262, 272, 262, 343, 343, 369, 383, 284, 284, 251, 272, 354, 265, 265, 304, 296, 276, 345, 287, 297, 526, 275, 321, 273, 327, 302, 349, 293, 374, 290, 328, 329, 340, 334, 312, 319, 339, 358, 361, 294, 276, 274, 312, 305, 328, 334, 332, 288, 348, 278, 258, 402, 337, 279, 330, 329, 292, 312, 358, 269, 370, 368, 358, 400, 349, 463, 445, 361, 387, 358, 368, 331, 348, 346, 445, 367, 362, 400, 299, 283, 327, 284, 340, 321, 364, 348, 286, 308, 340, 301, 315, 336, 286, 283, 271, 300, 281, 318, 489, 352, 341, 316, 282, 379, 349, 263, 353, 328, 353, 287, 298, 344, 319, 290, 292, 321, 352, 279, 336, 372, 342, 327, 264, 389, 335, 280, 275, 266, 310, 274, 259, 313, 354, 257, 295, 341, 269, 328, 323, 291, 433, 413, 338, 330, 286, 337, 296, 342, 304, 396, 473, 269, 298, 321, 269, 388, 357, 283, 330, 324, 370, 310, 355, 343, 283, 322, 333, 259, 285, 343, 275, 296, 339, 349, 330, 371, 380, 321, 289, 304, 343, 344, 316, 321, 289, 337, 279, 340, 331, 435, 351, 323, 385, 442, 388, 273, 306, 265, 345, 284, 302, 288, 377, 346, 536, 282, 285, 352, 289, 291, 293, 279, 279, 315, 371, 273, 396, 361, 296, 261, 266, 323, 286, 270, 264, 305, 275, 263, 297, 325, 279, 327, 279, 286, 265, 322, 265, 265, 259, 278, 281, 332, 268, 280, 266, 290, 371, 323, 254, 261, 285, 306, 337, 297, 317, 269, 300, 285, 291, 291, 273, 340, 248, 380, 266, 272, 361, 350, 286, 288, 408, 352, 340, 368, 344, 340, 287, 287, 379, 286, 304, 353, 428, 299, 351, 325, 346, 292, 369, 328, 301, 311, 348, 297, 337}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }

  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 30, MedianDifferenceThreshold: 20})
  assert.Equal(t, 0, len(degradations), "degradations: %v", degradations)
}

func TestComplexDistributionFromUnitTestWithDegradation(t *testing.T) {
  data := []int{101, 87, 88, 88, 90, 91, 97, 90, 101, 90, 89, 88, 97, 88, 92, 97, 86, 90, 93, 86, 143, 99, 93, 86, 87, 98, 87, 100, 100, 86, 99, 93, 95, 90, 88, 88, 98, 87, 89, 98, 88, 91, 88, 87, 101, 84, 97, 98, 95, 95, 92, 95, 91, 86, 95, 86, 96, 86, 86, 99, 98, 90, 101, 87, 100, 97, 88, 86, 95, 92, 98, 106, 97, 92, 94, 102, 86, 101, 86, 96, 101, 99, 86, 98, 98, 87, 102, 90, 108, 88, 97, 95, 94, 89, 87, 88, 86, 97, 94, 96, 97, 86, 101, 86, 90, 87, 87, 87, 88, 98, 94, 96, 95, 88, 102, 104, 102, 102, 87, 86, 87, 87, 100, 88, 96, 85, 88, 95, 86, 98, 98, 98, 98, 86, 92, 85, 96, 84, 93, 85, 87, 93, 86, 97, 88, 88, 88, 92, 97, 86, 93, 92, 90, 85, 100, 98, 98, 88, 94, 98, 93, 98, 88, 86, 100, 86, 102, 96, 92, 94, 92, 92, 118, 91, 89, 95, 102, 90, 97, 88, 95, 91, 87, 102, 98, 86, 87, 87, 104, 101, 97, 88, 99, 91, 85, 99, 93, 99, 100, 104, 101, 99, 100, 97, 87, 96, 82, 92, 87, 86, 97, 96, 86, 85, 89, 88, 89, 86, 92, 99, 92, 95, 94, 108, 89, 91, 94, 86, 91, 87, 86, 97, 87, 87, 87, 99, 89, 95, 97, 89, 88, 99, 86, 100, 100, 96, 86, 98, 87, 99, 101, 95, 85, 98, 85, 95, 86, 87, 89, 84, 86, 98, 88, 95, 98, 96, 99, 95, 93, 100, 97, 97, 88, 80, 87, 88, 88, 111, 92, 88, 88, 99, 97, 86, 85, 96, 95, 98, 90, 112, 87, 99, 99, 105, 92, 98, 92, 89, 97, 89, 96, 86, 96, 102, 91, 90, 86, 86, 97, 97, 96, 87, 86, 89, 106, 101, 91, 86, 90, 87, 85, 86, 101, 88, 98, 87, 88, 97, 97, 86, 88, 96, 96, 85, 101, 88, 95, 93, 92, 90, 88, 96, 84, 102, 87, 87, 108, 97, 107, 88, 89, 100, 93, 93, 95, 92, 101, 96, 98, 99, 97, 86, 98, 97, 96, 87, 98, 96, 97, 87, 97, 87, 94, 86, 87, 98, 86, 89, 94, 97, 99, 87, 97, 95, 96, 98, 101, 90, 96, 97, 97, 95, 97, 89, 83, 85, 93, 87, 86, 88, 102, 92, 84, 88, 86, 87, 96, 101, 96, 95, 96, 86, 92, 89, 89, 86, 96, 100, 87, 87, 97, 94, 103, 88, 84, 86, 95, 96, 85, 89, 91, 98, 85, 89, 97, 97, 84, 97, 91, 80, 88, 87, 96, 95, 93, 86, 150, 95, 86, 98, 98, 88, 101, 86, 99, 96, 87, 94, 90, 98, 85, 98, 95, 90, 97, 95, 96, 88, 103, 99, 95, 84, 86, 86, 99, 86, 98, 88, 91, 100, 96, 95, 92, 100, 99, 97, 85, 89, 96, 85, 97, 93, 98, 86, 84, 85, 87, 104, 102, 97, 97, 87, 96, 96, 96, 86, 86, 86, 85, 85, 89, 87, 97, 97, 86, 88, 92, 103, 88, 87, 83, 95, 101, 89, 92, 94, 86, 88, 90, 170, 98, 107, 85, 97, 90, 142, 95, 93, 99, 95, 83, 94, 97, 95, 97, 83, 89, 94, 98, 96, 96, 96, 98, 91, 95, 92, 91, 97, 100, 88, 99, 97, 97, 96, 98, 89, 95, 97, 98, 96, 96, 89, 97, 94, 80, 82, 89, 90, 92, 89, 86, 95, 96, 88, 97, 171, 85, 146, 98, 94, 94, 95, 86, 90, 87, 95, 96, 81, 89, 87, 87, 86, 86, 88, 93, 87, 86, 88, 87, 86, 89, 100, 98, 85, 88, 96, 89, 88, 103, 92, 87, 88, 90, 102, 99, 88, 87, 95, 98, 86, 86, 90, 86, 168, 86, 159, 99, 97, 103, 85, 91, 86, 87, 85, 95, 84, 105, 93, 88, 86, 94, 87, 85, 87, 85, 84, 97, 99, 86, 83, 91, 94, 92, 94, 98, 87, 96, 133, 122, 123, 124, 133, 136, 127, 132, 122, 135, 120, 125, 119, 129, 119, 124, 130, 121, 121, 136, 124, 133, 119, 130, 134, 130, 120, 120, 132, 124, 131, 138, 133, 132, 133, 139, 134, 131, 127, 131, 138, 135, 131, 123, 117, 121, 128, 124, 116, 122, 122, 121, 129, 141, 131, 136, 119, 118, 141, 130, 130, 131, 112, 121, 131, 122, 121, 119, 122, 133, 125, 118, 120, 139, 123, 117, 119, 131, 124, 119, 131, 121, 130, 122, 137, 122, 125, 120, 122, 121}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 30, MedianDifferenceThreshold: 20})
  assert.Equal(t, 1, len(degradations), "degradations: %v", degradations)
  assert.Equal(t, int64(667), degradations[0].timestamp)
}

func TestComplexDistributionFromUnitTestWithDegradation2(t *testing.T) {
  data := []int{608, 566, 556, 580, 565, 614, 580, 541, 576, 607, 565, 605, 574, 612, 585, 543, 591, 566, 563, 546, 548, 560, 527, 530, 577, 600, 579, 579, 587, 589, 611, 537, 555, 614, 538, 522, 562, 590, 609, 590, 570, 593, 578, 641, 532, 928, 533, 599, 568, 612, 613, 554, 535, 980, 603, 610, 610, 554, 544, 572, 530, 536, 566, 620, 595, 567, 605, 539, 605, 556, 558, 550, 594, 581, 606, 586, 608, 574, 611, 536, 548, 544, 543, 535, 549, 539, 614, 601, 577, 552, 572, 572, 604, 599, 540, 576, 527, 620, 613, 506, 525, 611, 606, 565, 559, 535, 582, 548, 591, 562, 533, 550, 607, 608, 539, 504, 514, 593, 627, 623, 571, 624, 606, 589, 589, 564, 586, 524, 613, 611, 556, 555, 541, 560, 562, 508, 611, 596, 573, 557, 585, 616, 575, 574, 529, 576, 591, 556, 569, 605, 605, 577, 538, 611, 588, 540, 556, 591, 523, 556, 542, 520, 589, 591, 597, 629, 601, 624, 579, 595, 592, 613, 595, 599, 590, 572, 588, 590, 593, 551, 580, 628, 576, 557, 529, 563, 632, 601, 591, 614, 602, 537, 565, 625, 575, 562, 616, 959, 596, 560, 513, 565, 581, 566, 602, 614, 568, 556, 588, 580, 563, 622, 563, 566, 567, 545, 515, 520, 573, 616, 563, 528, 587, 569, 565, 563, 518, 584, 545, 576, 514, 571, 532, 566, 581, 567, 562, 518, 497, 513, 567, 604, 585, 642, 575, 641, 618, 628, 545, 607, 578, 574, 754, 550, 610, 552, 525, 600, 539, 523, 548, 602, 550, 616, 561, 540, 568, 581, 582, 574, 561, 582, 513, 737, 571, 638, 600, 584, 576, 581, 565, 530, 617, 557, 540, 586, 528, 613, 579, 574, 600, 591, 552, 582, 591, 571, 569, 533, 548, 551, 551, 552, 590, 576, 608, 575, 623, 538, 545, 546, 556, 614, 555, 593, 594, 571, 598, 562, 529, 609, 598, 580, 594, 563, 572, 610, 520, 527, 568, 590, 576, 574, 591, 559, 523, 611, 584, 610, 584, 599, 543, 502, 592, 570, 550, 556, 566, 514, 567, 603, 562, 556, 630, 551, 563, 600, 564, 607, 566, 567, 597, 589, 519, 538, 540, 550, 631, 543, 540, 566, 592, 581, 570, 581, 553, 570, 551, 614, 556, 599, 584, 592, 541, 571, 559, 595, 571, 537, 595, 523, 565, 560, 598, 575, 600, 607, 567, 635, 542, 555, 555, 594, 531, 589, 546, 543, 573, 589, 564, 620, 597, 575, 541, 544, 636, 609, 577, 576, 505, 543, 592, 611, 577, 613, 529, 536, 634, 598, 620, 519, 628, 563, 636, 613, 537, 574, 597, 536, 531, 830, 832, 554, 586, 595, 538, 561, 619, 547, 544, 601, 596, 553, 544, 614, 606, 552, 571, 556, 542, 607, 546, 596, 569, 548, 552, 514, 628, 545, 617, 571, 543, 549, 612, 602, 608, 553, 583, 534, 587, 580, 542, 547, 556, 534, 562, 584, 591, 617, 541, 604, 605, 553, 566, 577, 605, 548, 563, 613, 556, 599, 499, 588, 615, 577, 571, 589, 604, 567, 610, 588, 564, 611, 598, 562, 599, 568, 553, 529, 557, 582, 547, 550, 596, 615, 554, 602, 595, 564, 911, 1041, 1086, 1059, 1070, 1087, 1071, 1028, 1039, 1065, 1035, 1067, 1062, 1052, 1016, 1046, 1074, 1071, 1089, 1006, 1087, 1065, 1080, 1152, 515, 548, 557, 541, 600, 529, 524, 595, 571, 576, 534, 533, 538, 586, 550, 541, 585, 539, 536, 585, 566, 550, 571, 533, 535, 596, 585, 614, 527, 560, 623, 537, 594, 658, 598, 617, 575, 608, 583, 602, 581, 547, 561, 578, 586, 570, 531, 526, 558, 598, 540, 571, 566, 599, 576, 583, 565, 595, 608, 614, 588, 602, 584, 534, 532, 551, 543, 563, 568, 574, 603, 622, 536, 535, 566, 573, 549, 598, 571, 557, 551, 528, 560, 555, 552, 571, 565, 599, 566, 568, 582, 595, 608, 567, 559, 529, 605, 606, 614, 541, 541, 566, 569, 607, 567, 575, 553, 611, 537, 569, 569, 590, 534, 604, 617, 577, 604, 602, 600, 564, 553, 587, 582, 623, 545, 532, 614, 570, 561, 606, 605, 586, 534, 600, 549, 604, 574, 573, 534, 563, 574, 595, 579, 562, 559, 595, 518, 611, 538, 599, 594, 597, 599, 568, 572, 591, 606, 545, 601, 565, 521, 602, 616, 593, 549, 558, 583, 506, 577, 599, 518, 539, 540, 588, 571, 618, 621, 573, 605, 550, 607, 569, 576, 651, 519, 601, 630, 566, 556, 584, 570, 528, 552, 539, 555, 589, 554, 598, 566, 639, 540, 570, 559, 629, 566, 622, 548, 543, 559, 604, 615, 530, 608, 611, 621, 607, 588, 561, 499, 560, 500, 531, 545, 597, 575, 572, 584, 613, 567, 549, 653, 566, 542, 532, 532, 540, 579, 593, 540, 556, 543, 550, 564, 545, 576, 566, 593, 612, 611, 565, 587, 556, 600, 623, 599, 536, 586, 570, 553, 542, 534, 527, 584, 559, 534, 549, 538, 586, 579, 528, 583, 548, 566, 563, 1068, 1072, 1050, 1048, 1043, 1030, 534, 564, 555, 572, 593, 555, 559, 601, 599, 539, 541, 531, 559, 544, 598, 585, 575, 540, 588, 537, 566, 575, 588, 539, 773, 768, 775, 750, 590, 786, 739, 603, 756, 798, 759, 786, 538, 598, 804, 568, 815, 784, 844, 758, 696, 766, 578, 820, 579, 633, 586, 616, 599, 735, 621, 831, 631, 591, 833, 745, 763, 877, 796, 583, 784, 549, 804, 753, 772, 696, 598, 593, 561, 596, 612, 810, 582, 567, 820, 798, 580, 627, 858, 758, 759, 736, 739, 907, 596, 588, 600, 595, 794, 603, 811, 780, 606, 1083, 1226, 1058, 653, 551, 613, 773, 746, 714, 770, 759, 809, 798, 770}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 30, MedianDifferenceThreshold: 20, EffectSizeThreshold: 0.7})
  assert.Equal(t, 1, len(degradations), "degradations: %v", degradations)
  assert.Equal(t, int64(856), degradations[0].timestamp)
}

func TestComplexDistributionFromUnitTestWithoutDegradation2(t *testing.T) {
  data := []int{1928, 1910, 1977, 2517, 2543, 1908, 1915, 2586, 2586, 1915, 2559, 1877, 2515, 1888, 1867, 2569, 2551, 2526, 1838, 2166, 1865, 1844, 2621, 1941, 2555, 1888, 2556, 1917, 2514, 1866, 1857, 1844, 2513, 2508, 2526, 2554, 2563, 1987, 2567, 1902, 2564, 1879, 2452, 1943, 1869, 2566, 1881, 1881, 1836, 2536, 1876, 2473, 1841, 1861, 2065, 2081, 2498, 1873, 2588, 2500, 1821, 2566, 2052, 1877, 2566, 1893, 2559, 2613, 2581, 1893, 2628, 2604, 2558, 2555, 2036, 1869, 2554, 1681, 2570, 2507, 2587, 1990, 2059, 2520, 1890, 2550, 2511, 1923, 2512, 2513, 2610, 2524, 2638, 2517, 2518, 2519, 2521, 1899, 1866, 2526, 2536, 1873, 1917, 2519, 2460, 1904, 2506, 2567, 1921, 2570, 1984, 2638, 1874, 1835, 2546, 2512, 2611, 1929, 2033, 1905, 2458, 2512, 1915, 1980, 1930, 1990, 1935, 1869, 2044, 2513, 1872, 1895, 1876, 1890, 2556, 2081, 2558, 1938, 1872, 1878, 1874, 2531, 2517, 2541, 2056, 1878, 1874, 2516, 1941, 2564, 2509, 1874, 1907, 1931, 1870, 1870, 1836, 2600, 1861, 1890, 2486, 2057, 2049, 1902, 2581, 1874, 2519, 2484, 1845, 2542, 1903, 1888, 2600, 2021, 2569, 1987, 2001, 2662, 1939, 1957, 2543, 1830, 2007, 2608, 1938, 1866, 2574, 2638, 2519, 2607, 2519, 1906, 2590, 2052, 2531, 2574, 1937, 2550, 1890, 1870, 2626, 2550, 1890, 2565, 2559, 1867, 2561, 2014, 2584, 2502, 1872, 1882, 1880, 1920, 1839, 1879, 1872, 2536, 2525, 1995, 1936, 1882, 1873, 2597, 2558, 1890, 2693, 1837, 2554, 2514, 1873, 2616, 1938, 2557, 1935, 1828, 2094, 2583, 1916, 1919, 1893, 2522, 1870, 2534, 2023, 1935, 2557, 2458, 1892, 2506, 2600, 2028, 2536, 1923, 2591, 2624, 2568, 1871, 1933, 1920, 2514, 2499, 2528, 1858, 2538, 1877, 2056, 1989, 1872, 1893, 2510, 2528, 2508, 1837, 2581, 2025, 2529, 1850, 2537, 1865, 2550, 1850, 1988, 1874, 2606, 1868, 2569, 1952, 1871, 2566, 1969, 2555, 2002, 1853, 2000, 2509, 2507, 2556, 2638, 2580, 2019, 1918, 2523, 1881, 2511, 1880, 1934, 1854, 1878, 2649, 2581, 1885, 1961, 1933, 2567, 2579, 1902, 2509, 2028, 1927, 1927, 2529, 2549, 2552, 2509, 2515, 1939, 1876, 1996, 1916, 2040, 2043, 1877, 1879, 2578, 2079, 2516, 1878, 1914, 2558, 1939, 2055, 1869, 2558, 2570, 2573, 1875, 2557, 2511, 2551, 1935, 1876, 2570, 2528, 2607, 2075, 2458, 2615, 2010, 2519, 1928, 1920, 2588, 2564, 1920, 2485, 2562, 2596, 2602, 2567, 2619, 1876, 2586, 2564, 2559, 1907, 1926, 1956, 2566, 1966, 2547, 2067, 1871, 1872, 2001, 1884, 2547, 1875, 1917, 1991, 2625, 2519, 2506, 1899, 2520, 2011, 2016, 2558, 1832, 1849, 1912, 2514, 1902, 2568, 2507, 2526, 2522, 2558, 2589, 1904, 1937, 2574, 1866, 1877, 2520, 1872, 2575, 2594, 2505, 2031, 1874, 2510, 1892, 2634, 1850, 1875, 2557, 1900, 2618, 2540, 1805, 2549, 2559, 2537, 1884, 2518, 2556, 2612, 2559, 1877, 2517, 2520, 1906, 2560, 2507, 2599, 1871, 2558, 2538, 1855, 1934, 1845, 2524, 2053, 2523, 1914, 1986, 1873, 1879, 1917, 2594, 2565, 1887, 1881, 1935, 1911, 1905, 2524, 2539, 2539, 2534, 2024, 1915, 2063, 2016, 2572, 1877, 1921, 1930, 2514, 1793, 2026, 1871, 2571, 2546, 2050, 1868, 2577, 1905, 2559, 2552, 2131, 2050, 1846, 1797, 1750, 1727, 1727, 2162, 2026, 1983, 1940, 1740, 2199, 2230, 2089, 2125, 2161, 1876, 2064, 2235, 1939, 1902, 1824, 2131, 1741, 3727, 3724, 2777, 2750, 2100, 2215, 1871, 2618, 1870, 2522, 1891, 2574, 2508, 2579, 2593, 1875, 2615, 2565, 2535, 1957, 1841, 2562, 2573, 2593, 2562, 2553, 1992, 2561, 2554, 2582, 2565, 1913, 2459, 1945, 1893, 2519, 1993, 2566, 1889, 1869, 2517, 1992, 2553, 2534, 2508, 1875, 2590, 1927, 2023, 1980, 1924, 2567, 1946, 1870, 2569, 2539, 1640, 1866, 2507, 1879, 1872, 2581, 1847, 1899, 1998, 1931, 2011, 1889, 1898, 1870, 1986, 1880, 2731, 2627, 2570, 2577, 1908, 1920, 2014, 1916, 1877, 2534, 2509, 2523, 1918, 2577, 2548, 1933, 1916, 2532, 2064, 1838, 1839, 2517, 1873, 1629, 2583, 1883, 2599, 1914, 2568, 1834, 2563, 2513, 1916, 2510, 2587, 2507, 1929, 1910, 1836, 1865, 2523, 2524, 2552, 2567, 1912, 2561, 2573, 2527, 2517, 2015, 2704, 1821, 2558, 1869, 2639, 1894, 2578, 1880, 1891, 2597, 1849, 2520, 1907, 2567, 1852, 1822, 1943, 2617, 1850, 2535, 1859, 2515, 1881, 1703, 1935, 2527, 2570, 2520, 1838, 1875, 1942, 2568, 2552, 1870, 2673, 2603, 1964, 2525, 2499, 2579, 2568, 2518, 2520, 1874, 2008, 1901, 1911, 2458, 2567, 2033, 2562, 2687, 1941, 2533, 1827, 2547, 1901, 2530, 2519, 2077, 2007, 2562, 2559, 2510, 1920, 1987, 2511, 1893, 1865, 2579, 1912, 2570, 1962, 2514, 2528, 2481, 2557, 1875, 1623, 1877, 2585, 2579, 2048, 2616, 1903, 1921, 2548, 2516, 2531, 1895, 1922, 2532, 1885, 2519, 2461, 2587, 1875, 2537, 1909, 1942, 2579, 1925, 1875, 1901, 2521, 2590, 2633, 2467, 1936, 2546, 2504, 2504, 2520, 1926, 2637, 2591, 1917, 1671, 1624, 1878, 1923, 2580, 1923, 1907, 2525, 1867, 2516, 2517, 1938, 2556, 2529, 1889, 2555, 1830, 2572, 2576, 2594, 2581, 1900, 1883, 2516, 2507, 1894, 2020, 2517, 2612, 2596, 2514, 1842, 2509, 2012, 1870, 2597, 2510, 2023, 1862, 1998, 1867, 2586, 2543, 1831, 2557, 1903, 2493, 1838, 2551, 1874, 1864, 2601, 2515, 1852, 1936, 2527, 1885, 1850, 2508, 2584, 2545, 2512, 2515, 2516, 2552, 1867, 2106, 2025, 2186, 2171, 1852, 2086, 2233, 3727, 2764, 2733, 2741, 3715, 3734, 2574, 2532, 2514, 2565, 2557, 2572, 1926, 2594, 1879, 1991, 2558, 1906, 1884, 1868, 2516, 2561, 1692, 2585, 2591, 2600, 2037, 2021, 2518, 2526, 1830, 2517, 2557, 1892, 1866, 1892, 1868, 2517, 1918, 2501, 1876, 2535, 1911, 1940, 1873, 1903, 2570, 2566, 1924, 2528, 1837, 2496, 1924, 2555, 2587, 1999, 2552, 2546, 2583}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 30, MedianDifferenceThreshold: 20, EffectSizeThreshold: 0.2})
  assert.Equal(t, 0, len(degradations), "degradations: %v", degradations)
}

func TestComplexDistributionFromUnitTestWithoutDegradation3(t *testing.T) {
  data := []int{662, 627, 297, 412, 608, 263, 284, 294, 295, 153, 635, 423, 720, 509, 159, 169, 479, 403, 572, 225, 642, 598, 277, 294, 170, 645, 187, 187, 693, 176, 695, 283, 270, 524, 605, 169, 302, 575, 217, 307, 301, 639, 173, 698, 186, 693, 300, 672, 293, 674, 170, 192, 154, 652, 686, 677, 677, 306, 627, 368, 602, 291, 534, 178, 296, 627, 285, 695, 162, 152, 610, 655, 525, 262, 339, 622, 605, 174, 321, 312, 294, 274, 279, 294, 282, 230, 590, 640, 460, 296, 313, 670, 521, 298, 152, 456, 214, 291, 280, 161, 352, 288, 620, 160, 199, 304, 153, 318, 169, 281, 694, 310, 637, 300, 280, 172, 309, 151, 177, 154, 647, 309, 334, 614, 281, 319, 665, 673, 290, 292, 296, 279, 395, 156, 602, 560, 612, 373, 310, 306, 289, 154, 334, 285, 284, 155, 313, 294, 570, 616, 280, 632, 553, 231, 631, 301, 288, 289, 157, 268, 644, 591, 511, 242, 197, 209, 591, 224, 315, 705, 616, 187, 558, 163, 619, 204, 511, 308, 198, 178, 180, 616, 285, 289, 212, 174, 448, 642, 695, 280, 597, 586, 298, 623, 619, 875, 316, 643, 291, 627, 693, 299, 184, 210, 460, 432, 158, 678, 437, 483, 603, 673, 290, 173, 270, 281, 285, 624, 156, 286, 286, 640, 283, 301, 296, 215, 483, 610, 558, 281, 295, 288, 627, 289, 156, 287, 277, 281, 355, 637, 225, 285, 620, 644, 280, 601, 299, 640, 614, 222, 495, 280, 399, 295, 283, 537, 286, 300, 327, 353, 671, 692, 565, 255, 683, 633, 673, 152, 281, 569, 285, 697, 637, 638, 688, 463, 636, 642, 288, 318, 357, 348, 293, 684, 299, 578, 291, 551, 325, 305, 563, 192, 628, 206, 553, 209, 164, 148, 313, 279, 604, 602, 319, 434, 669, 169, 188, 276, 155, 552, 284, 318, 703, 604, 467, 285, 678, 296, 657, 597, 607, 170, 275, 347, 278, 283, 285, 594, 645, 286, 599, 170, 280, 436, 564, 690, 632, 591, 289, 602, 684, 701, 308, 324, 393, 282, 663, 295, 547, 570, 488, 187, 512, 245, 569, 163, 623, 343, 645, 634, 190, 283, 191, 153, 655, 290, 288, 614, 144, 615, 523, 291, 226, 162, 284, 469, 298, 297, 525, 530, 275, 299, 290, 448, 694, 297, 341, 556, 531, 187, 632, 628, 288, 586, 286, 293, 630, 275, 286, 543, 274, 581, 573, 278, 465, 294, 289, 287, 328, 319, 604, 293, 297, 333, 504, 317, 150, 305, 689, 654, 233, 152, 292, 158, 607, 156, 683, 295, 151, 171, 339, 352, 655, 605, 433, 281, 150, 771, 149, 285, 599, 330, 585, 287, 717, 624, 291, 610, 280, 314, 289, 235, 646, 668, 611, 477, 160, 465, 558, 350, 150, 295, 341, 299, 670, 285, 167, 299, 170, 286, 293, 318, 474, 284, 285, 152, 295, 330, 283, 283, 293, 284, 280, 646, 571, 162, 283, 590, 640, 289, 668, 359, 630, 172, 589, 288, 587, 160, 278, 559, 592, 315, 298, 154, 312, 599, 192, 302, 604, 359, 327, 241, 321, 287, 326, 277, 283, 311, 295, 243, 424, 352, 677, 412, 203, 147, 194, 721, 726, 755, 610, 763, 183, 618, 479, 528, 746, 754, 763, 468, 700, 593, 718, 599, 587, 296, 756, 634, 697, 632, 220, 745, 272, 389, 348, 303, 553, 163, 334, 432, 598, 683, 286, 281, 287, 651, 286, 642, 684, 644, 185, 425, 469, 604, 490, 286, 182, 305, 499, 280, 595, 287, 204, 304, 687, 478, 300, 308, 216, 584, 176, 288, 666, 395, 591, 485, 494, 656, 528, 280, 616, 296, 168, 609, 289, 152, 630, 300, 689, 630, 279, 437, 327, 357, 414, 283, 662, 277, 296, 298, 635, 290, 695, 315, 179, 154, 464, 677, 284, 300, 510, 287, 283, 282, 284, 407, 171, 609, 285, 276, 657, 338, 534, 619, 165, 151, 300, 279, 613, 494, 223, 287, 279, 145, 283, 300, 191, 297, 311, 644, 298, 577, 577, 239, 287, 170, 157, 147, 293, 464, 389, 521, 308, 552, 315, 479, 277, 294, 670, 191, 412, 392, 402, 284, 294, 285, 272, 525, 646, 315, 459, 284, 620, 324, 197, 679, 571, 650, 292, 565, 169, 165, 554, 649, 344, 359, 552, 546, 626, 278, 543, 348, 286, 182, 308, 633, 273, 167, 640, 283, 279, 675, 147, 358, 161, 284, 550, 455, 651, 378, 176, 473, 549, 612, 310, 377, 588, 157, 371, 635, 292, 304, 670, 358, 301, 148, 287, 610, 285, 702, 277, 700, 307, 639, 556, 607, 479, 348, 291, 313, 286, 173, 553, 274, 364, 652, 717, 661, 280, 453, 280, 303, 297, 300, 543, 659, 443, 314, 289, 184, 515, 298, 290, 475, 599, 276, 160, 304, 155, 661, 286, 591, 300, 615, 440, 295, 619, 514, 590, 579, 333, 296, 298, 498, 512, 629, 664, 154, 181, 675, 628, 156, 171, 583, 672, 151, 282, 160, 596, 321, 299, 294, 604, 304, 163, 571, 739, 762, 602, 499, 201, 706, 643, 707, 690, 650, 713, 365, 295, 549, 571, 203, 309, 485, 304, 703, 291, 608, 594, 332, 321, 329, 682, 530, 613, 291, 279, 627, 563, 447, 524, 637, 298, 444, 310, 538, 629, 649, 684, 300, 302, 296, 638, 652, 288, 661, 346, 651, 636, 298, 300, 295, 304, 512, 303, 535, 681, 722, 526, 586, 321, 706, 317, 685, 635, 299, 291, 529, 589, 621, 364, 468, 611, 441, 302, 299, 292, 591, 513, 286, 200, 245, 304, 599, 698, 569, 306, 457, 231, 464, 320, 309, 300, 307, 704, 277, 297, 309, 531, 327, 575, 563, 311, 665, 703, 674, 520, 645, 677, 673, 361, 627, 696, 309, 584, 553, 531, 602, 626, 634}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := detectDegradations(data, builds, times, AnalysisSettings{MinimumSegmentLength: 30, MedianDifferenceThreshold: 20, EffectSizeThreshold: 1.2})
  assert.Equal(t, 0, len(degradations), "degradations: %v", degradations)
}
