package template

func init() {
	CommonProjectFiles[".gitlab-ci.yml"] = `# Created by template.
stages:
  - test
  - build
  - deploy

test:
  stage: test
  image: golang:1.15-alpine
  script:
    - export GO111MODULE=on
    - CGO_ENABLED=0 go test ./... -p 1 -v -coverprofile=coverage.out
    - go tool cover -func=coverage.out
  only:
    - develop
    - merge_requests
    - tags
  except:
    - triggers
    - apis
  coverage: '/^coverage:(\s)+(\d+\.\d+%)/'

build:
  stage: build
  script:
    - ~/ci-script/build_image.sh
  only:
    refs:
      - develop
	  - tags
  except:
    - triggers
	- apis
  tags:
    - image-builder-cn

k8s-staging:
  stage: deploy
  script:
    - . ~/ci-script/deploy.sh staging
  only:
    - develop
  except:
    - triggers
	- apis
  tags:
    - deployer-cn

k8s-pre-release:
  stage: deploy
  script:
    - . ~/ci-script/deploy.sh pre-release
  only:
    - tags
  tags:
    - deployer-us
  when: manual

k8s-production:
  stage: deploy
  script:
    - . ~/ci-script/deploy.sh production
  only:
    - tags
  tags:
    - deployer-us
  when: manual
`
}
