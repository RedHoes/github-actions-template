plugins {
    kotlin("jvm") version "1.9.20"
    application
}

group = "com.example"
version = "1.0.0"

java {
    toolchain {
        languageVersion.set(JavaLanguageVersion.of(17))
    }
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("org.jetbrains.kotlin:kotlin-stdlib")
    testImplementation("org.junit.jupiter:junit-jupiter:5.10.0")
    testImplementation("org.jetbrains.kotlin:kotlin-test")
}

tasks.test {
    useJUnitPlatform()
}

tasks.jar {
    archiveBaseName.set("backend-kotlin")
    archiveVersion.set(project.version.toString())
    manifest {
        attributes(mapOf("Main-Class" to "com.example.MainKt"))
    }
}

application {
    mainClass.set("com.example.MainKt")
}
