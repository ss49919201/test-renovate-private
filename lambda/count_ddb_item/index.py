from pprint import pprint
import boto3
from boto3.dynamodb.conditions import Key
import os


def count_books(date):
    dynamodb = boto3.resource('dynamodb')

    table = dynamodb.Table('books')

    response = table.query(
        KeyConditionExpression=Key('date').eq(date),
        Select='COUNT',
    )
    return response['Count']


def publish_message(count):
    arn = os.environ['SNS_TOPIC_ARN']

    msg = str(count) + '件の本が登録されました'
    subject = '本登録情報'
    client = boto3.client('sns')

    request = {
        'TopicArn': arn,
        'Message': msg,
        'Subject': subject
    }

    response = client.publish(**request)
    return response


if __name__ == '__main__':
    count = count_books('2021-12-22')
    if count >= 0:
        response = publish_message(count)
        pprint(response)
