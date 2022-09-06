DROP TABLE IF EXISTS v5FlowHeader;
DROP TABLE IF EXISTS v5FlowBody;

CREATE TABLE v5FlowHeader (
                    id         INT AUTO_INCREMENT NOT NULL,
                    PRIMARY KEY (`id`),
                    version               VARCHAR(128) NOT NULL,
                    count                 VARCHAR(255) NOT NULL,
                    sysUptime             DECIMAL(5,2) NOT NULL,
                    unixSecs              DECIMAL(5,2) NOT NULL,
                    unixNsecs             DECIMAL(5,2) NOT NULL,
                    flowSequence          DECIMAL(5,2) NOT NULL,
                    engineType            DECIMAL(5,2) NOT NULL,
                    engineId              DECIMAL(5,2) NOT NULL,
                    samplingInterval      DECIMAL(5,2) NOT NULL,
                    srcaddr               INT NOT NULL,
                    dstaddr               INT NOT NULL,
                    nexthop               INT NOT NULL,
                    input                 SMALLINT,
                    output                SMALLINT,
                    dPkts                 INT NOT NULL ,
                    dOctets               INT NOT NULL,
                    First                 INT NOT NULL ,
	                Last                  INT NOT NULL,
	                srcport               SMALLINT,
	                dstport               SMALLINT,
	                pad1                  SMALLINT,
	                tcpFlags              SMALLINT,
	                prot                  SMALLINT,
	                tos                   SMALLINT,
	                srcAs                 INT,
	                dstAs                 INT,
	                srcMask               SMALLINT,
	                dstMask               SMALLINT,
	                pad2                  INT
};

CREATE TABLE netflow7 (
                          id         INT AUTO_INCREMENT NOT NULL,
                          title      VARCHAR(128) NOT NULL,
                          artist     VARCHAR(255) NOT NULL,
                          price      DECIMAL(5,2) NOT NULL,
                          PRIMARY KEY (`id`)
);

CREATE TABLE netflow9 (
                          id         INT AUTO_INCREMENT NOT NULL,
                          title      VARCHAR(128) NOT NULL,
                          artist     VARCHAR(255) NOT NULL,
                          price      DECIMAL(5,2) NOT NULL,
                          PRIMARY KEY (`id`)
);

INSERT INTO netflow5
(title, artist, price)
VALUES
    ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);

INSERT INTO netflow7
(title, artist, price)
VALUES
    ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);

INSERT INTO netflow9
(title, artist, price)
VALUES
    ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);