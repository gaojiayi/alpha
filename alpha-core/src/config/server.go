package config

const PORT = ":9090"
// 因为涉及到图像的探测  可能会耗时   暂时设置为1分钟
const READ_TIME_OUT = 60
const WRITE_TIME_OUT = 60
const MAX_HEADER_BYTE = 1 << 20

