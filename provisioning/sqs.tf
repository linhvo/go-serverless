resource "aws_sqs_queue" "terraform_queue_deadletter" {
  name = "${var.queue_name}-dead-letter"
  delay_seconds = 90
  max_message_size = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
  tags                      = "${var.default_tags}"
}
resource "aws_sqs_queue" "terraform_queue" {
  name = "${var.queue_name}"
  delay_seconds = 90
  max_message_size = 2048
  message_retention_seconds = "${var.message_retention_seconds}"
  receive_wait_time_seconds = 10
  redrive_policy = "{\"deadLetterTargetArn\":\"${aws_sqs_queue.terraform_queue_deadletter.arn}\",\"maxReceiveCount\":${var.maxReceiveCountBeforeDLQ}}"
  tags                      = "${var.default_tags}"
}

//resource "aws_sqs_queue_policy" "${var.queue_name}_policy" {
//  queue_url = "${aws_sqs_queue.terraform_queue.id}"
//
//  policy = <<POLICY
//{
//  "Version": "2012-10-17",
//  "Id": "sqspolicy",
//  "Statement": [
//    {
//      "Effect": "Allow",
//      "Principal": "*",
//      "Action": [
//        "sqs:ReceiveMessage",
//        "sqs:DeleteMessage",
//        "sqs:GetQueueAttributes"
//      ],
//      "Resource": "${aws_sqs_queue.terraform_queue.arn}",
//      "Condition": {
//        "ArnEquals": {
//          "aws:SourceArn": "${aws_lambda_function.testgo-dev-hello.arn}"
//        }
//      }
//    }
//  ]
//}
//POLICY
//}
