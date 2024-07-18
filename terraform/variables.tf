variable "workload" {
  description = "ワークロード名"
  type        = string
}

variable "location" {
  description = "リソースが作成されるリージョン。既定は東日本"
  type        = list(string)
  default     = ["japaneast", "jpe"]
}

variable "tags" {
  description = "リソースに付加されるタグ。"
  type        = map(string)
}